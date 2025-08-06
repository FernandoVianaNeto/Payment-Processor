package payment_usecase

import (
	"context"
	"encoding/json"
	"errors"
	"payment-gateway/internal/domain/adapters/queue"
	"payment-gateway/internal/domain/dto"
	"payment-gateway/internal/domain/entity"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"
	"time"
)

type CreatePaymentRequestUsecase struct {
	PaymentRepository domain_repository.PaymentRepositoryInterface
	Queue             queue.Client
}

func NewCreatePaymentRequestUsecase(
	paymentRepository domain_repository.PaymentRepositoryInterface,
	queue queue.Client,
) domain_payment_usecase.CreatePaymentRequestUsecaseInterface {
	return &CreatePaymentRequestUsecase{
		PaymentRepository: paymentRepository,
		Queue:             queue,
	}
}

func (u *CreatePaymentRequestUsecase) Execute(ctx context.Context, dto dto.CreatePaymentDto) error {
	paymentAlreadyProcessed := u.PaymentRepository.AlreadyAdded(ctx, dto.CorrelationId)
	if paymentAlreadyProcessed {
		return errors.New("payment already processed")
	}

	now := time.Now().UTC()
	requestedAt := now.Format(time.RFC3339)

	message := entity.NewPayment(dto.CorrelationId, dto.Amount, requestedAt)
	messageByte, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = u.Queue.Publish("payment_requests", messageByte)
	return err
}
