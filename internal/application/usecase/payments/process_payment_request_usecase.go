package payment_usecase

import (
	"context"
	"encoding/json"
	"log"
	configs "payment-gateway/cmd/config"
	"payment-gateway/internal/domain/adapters/processors"
	"payment-gateway/internal/domain/adapters/queue"
	"payment-gateway/internal/domain/dto"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_response "payment-gateway/internal/domain/response"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"
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
	log.Println("PROCESSING DEFAULT MESSAGE WITH CORRELATION ID: ", message.CorrelationId)

	defaultHealthCheckResponse, err := u.ProcessPaymentDefaultAdapter.IsLive(ctx)

	log.Println("DEFAULT HEALTH CHECK RESPONSE: ", defaultHealthCheckResponse)

	err = u.failHealthCheck(defaultHealthCheckResponse, err, message)

	if err != nil {
		return err
	}

	input := processors.ProcessorClientInput{
		CorrelationId: message.CorrelationId,
		Amount:        message.Amount,
		RequestedAt:   message.RequestedAt,
	}

	log.Println("EXECUTING DEFAULT PAYMENT: ", defaultHealthCheckResponse)

	err = u.ProcessPaymentDefaultAdapter.ExecutePayment(ctx, input)

	if err != nil {
		newData, err := retryMessage(message)

		if err != nil {
			log.Println("error processing default payment. Sending an retry message")

			u.Queue.Publish(configs.NatsCfg.PaymentRequestsQueue, newData)

			return err
		}

		return err
	}

	err = u.PaymentRepository.Create(ctx, dto.CreatePaymentDto{
		CorrelationId:        message.CorrelationId,
		Amount:               message.Amount,
		RequestedAt:          message.RequestedAt,
		TransactionProcessor: "default",
	})

	if err != nil {
		log.Println("could not save payment into database when processing default payment. Sending retry message to database fallback...", err)

		return err
	}

	log.Println("default payment processed successfully", err)
	return err
}

func (u *ProcessPaymentRequestUsecase) processFallbackMessage(ctx context.Context, message dto.ProcessPaymentRequestDto) error {
	log.Println("PROCESSING FALLBACK MESSAGE WITH CORRELATION ID: ", message.CorrelationId)
	fallbackHealthCheckResponse, err := u.ProcessPaymentFallbackAdapter.IsLive(ctx)

	log.Println("FALLBACK HEALTH CHECK RESPONSE: ", fallbackHealthCheckResponse)

	err = u.failHealthCheck(fallbackHealthCheckResponse, err, message)

	if err != nil {
		return err
	}

	err = u.ProcessPaymentFallbackAdapter.ExecutePayment(ctx, processors.ProcessorClientInput{
		CorrelationId: message.CorrelationId,
		Amount:        message.Amount,
		RequestedAt:   message.RequestedAt,
	})

	if err != nil {
		newData, err := retryMessage(message)

		if err != nil {
			log.Println("error processing fallback payment. Sending an retry message")

			u.Queue.Publish(configs.NatsCfg.PaymentRequestsQueue, newData)

			return err
		}
	}

	err = u.PaymentRepository.Create(ctx, dto.CreatePaymentDto{
		CorrelationId:        message.CorrelationId,
		Amount:               message.Amount,
		RequestedAt:          message.RequestedAt,
		TransactionProcessor: "fallback",
	})

	if err != nil {
		log.Println("could not save payment into database when processing default payment. Sending retry message to database fallback...")
		return err
	}

	log.Println("fallback payment processed successfully")
	return err
}

func (u *ProcessPaymentRequestUsecase) failHealthCheck(healthCheckResponse *domain_response.HealthCheckResponse, err error, message dto.ProcessPaymentRequestDto) error {
	isFailing := &healthCheckResponse.Failing

	if err != nil || *isFailing {
		newData, err := retryMessage(message)

		if err != nil {
			log.Println("error getting health check. Sending an retry message")

			u.Queue.Publish(configs.NatsCfg.PaymentRequestsQueue, newData)

			return err
		}
	}

	return nil
}

func retryMessage(message dto.ProcessPaymentRequestDto) ([]byte, error) {
	newData, err := json.Marshal(dto.ProcessPaymentRequestDto{
		CorrelationId: message.CorrelationId,
		Amount:        message.Amount,
		RequestedAt:   message.RequestedAt,
		Retries:       message.Retries + 1,
	})

	if err != nil {
		log.Println("could not marshall new message: ", err)
		return []byte{}, err
	}

	return newData, err
}
