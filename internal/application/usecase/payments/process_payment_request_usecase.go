package payment_usecase

import (
	"context"
	"log"
	configs "payment-gateway/cmd/config"
	"payment-gateway/internal/domain/adapters/processors"
	"payment-gateway/internal/domain/adapters/queue"
	"payment-gateway/internal/domain/dto"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"
	"time"

	jsoniter "github.com/json-iterator/go"
)

const MAX_AMOUNT_OF_RETRIES = 5

type ProcessPaymentRequestUsecase struct {
	PaymentRepository             domain_repository.PaymentRepositoryInterface
	ProcessPaymentDefaultAdapter  processors.ProcessorsClientInterface
	ProcessPaymentFallbackAdapter processors.ProcessorsClientInterface
	Queue                         queue.Client
}

func NewProcessPaymentRequestUsecase(
	paymentRepository domain_repository.PaymentRepositoryInterface,
	processPaymentDefaultAdapter processors.ProcessorsClientInterface,
	processPaymentFallbackAdapter processors.ProcessorsClientInterface,
	queue queue.Client,
) domain_payment_usecase.ProcessPaymentRequestUsecaseInterface {
	return &ProcessPaymentRequestUsecase{
		PaymentRepository:             paymentRepository,
		ProcessPaymentDefaultAdapter:  processPaymentDefaultAdapter,
		ProcessPaymentFallbackAdapter: processPaymentFallbackAdapter,
		Queue:                         queue,
	}
}

func (u *ProcessPaymentRequestUsecase) Execute(ctx context.Context, data []byte) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var message dto.ProcessPaymentRequestDto

	err := json.Unmarshal(data, &message)

	if err != nil {
		log.Println("Error on unmarshalling message: ", err)
		return err
	}

	if message.Retries >= MAX_AMOUNT_OF_RETRIES {
		err = u.processFallbackMessage(ctx, message)
		return err
	}

	err = u.processDefaultMessage(ctx, message)

	return err
}

func (u *ProcessPaymentRequestUsecase) processDefaultMessage(ctx context.Context, message dto.ProcessPaymentRequestDto) error {
	now := time.Now().UTC()
	requestedAt := now.Format(time.RFC3339)

	input := processors.ProcessorClientInput{
		CorrelationId: message.CorrelationId,
		Amount:        message.Amount,
		RequestedAt:   requestedAt,
	}
	err := u.ProcessPaymentDefaultAdapter.ExecutePayment(ctx, input)

	if err == nil {
		err := u.PaymentRepository.Create(ctx, dto.CreatePaymentDto{
			CorrelationId:        message.CorrelationId,
			Amount:               message.Amount,
			RequestedAt:          message.RequestedAt,
			TransactionProcessor: "default",
		})

		return err
	}
	err = u.retryMessage(message)

	return err
}

func (u *ProcessPaymentRequestUsecase) processFallbackMessage(ctx context.Context, message dto.ProcessPaymentRequestDto) error {
	err := u.PaymentRepository.Create(ctx, dto.CreatePaymentDto{
		CorrelationId:        message.CorrelationId,
		Amount:               message.Amount,
		RequestedAt:          message.RequestedAt,
		TransactionProcessor: "fallback",
	})
	if err != nil {
		err := u.retryMessage(message)
		return err
	}
	err = u.ProcessPaymentFallbackAdapter.ExecutePayment(ctx, processors.ProcessorClientInput{
		CorrelationId: message.CorrelationId,
		Amount:        message.Amount,
		RequestedAt:   message.RequestedAt,
	})
	if err != nil {
		u.PaymentRepository.Delete(ctx, message.CorrelationId)
		err := u.retryMessage(message)
		return err
	}
	return err
}

// func (u *ProcessPaymentRequestUsecase) failHealthCheck(healthCheckResponse *domain_response.HealthCheckResponse, err error, message dto.ProcessPaymentRequestDto) error {
// 	isFailing := &healthCheckResponse.Failing

// 	if err != nil || healthCheckResponse == nil || *isFailing {
// 		newData, err := retryMessage(message)

// 		if err != nil {
// 			log.Println("error getting health check. Sending an retry message")

// 			u.Queue.Publish(configs.NatsCfg.PaymentRequestsQueue, newData)

// 			return err
// 		}
// 	}

// 	return nil
// }

func (u *ProcessPaymentRequestUsecase) retryMessage(message dto.ProcessPaymentRequestDto) error {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	newData, err := json.Marshal(dto.ProcessPaymentRequestDto{
		CorrelationId: message.CorrelationId,
		Amount:        message.Amount,
		RequestedAt:   message.RequestedAt,
		Retries:       message.Retries + 1,
	})

	if err != nil {
		log.Println("could not marshall new message: ", err)
		return err
	}

	err = u.Queue.Publish(configs.NatsCfg.PaymentRequestsQueue, newData)

	return err
}
