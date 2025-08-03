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
		fallbackHealthCheckResponse, err := u.ProcessPaymentFallbackAdapter.IsLive(ctx)

		err = u.failHealthCheck(fallbackHealthCheckResponse, err, message)

		if err != nil {
			return err
		}

		return err
	}

	defaultHealthCheckResponse, err := u.ProcessPaymentDefaultAdapter.IsLive(ctx)

	err = u.failHealthCheck(defaultHealthCheckResponse, err, message)

	if err != nil {
		return err
	}

	

	return err
}

func (u *ProcessPaymentRequestUsecase) failHealthCheck(healthCheckResponse *domain_response.HealthCheckResponse, err error, message dto.ProcessPaymentRequestDto) error {
	isFailing := &healthCheckResponse.Failing

	if err != nil || *isFailing {
		newData, err := sendRetryMessage(message)

		if err != nil {
			return err
		}

		u.Queue.Publish(configs.NatsCfg.PaymentRequestsQueue, newData)
	}

	return nil
}

func sendRetryMessage(message dto.ProcessPaymentRequestDto) ([]byte, error) {
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
