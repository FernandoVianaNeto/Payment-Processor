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
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"
	"time"
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

	err = u.processDefaultMessage(ctx, message)

	return err
}

func (u *ProcessPaymentRequestUsecase) processDefaultMessage(ctx context.Context, message dto.ProcessPaymentRequestDto) error {
	input := processors.ProcessorClientInput{
		CorrelationId: message.CorrelationId,
		Amount:        message.Amount,
		RequestedAt:   time.Now().UTC().Format(time.RFC3339),
	}

	err := u.ProcessPaymentDefaultAdapter.ExecutePayment(ctx, input)

	if err != nil {
		err := u.retryMessage(message)

		return err
	}

	err = u.PaymentRepository.Create(ctx, dto.CreatePaymentDto{
		CorrelationId:        message.CorrelationId,
		Amount:               message.Amount,
		RequestedAt:          message.RequestedAt,
		TransactionProcessor: "default",
	})

	return err
}

func (u *ProcessPaymentRequestUsecase) retryMessage(message dto.ProcessPaymentRequestDto) error {
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

	return u.Queue.Publish(configs.NatsCfg.PaymentRequestsQueue, newData)
}
