package payment_usecase

import (
	"context"
	"errors"
	"payment-gateway/internal/domain/adapters/messaging"
	"payment-gateway/internal/domain/dto"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"
)

type CreatePaymentUsecase struct {
	PaymentRepository domain_repository.PaymentRepositoryInterface
	Queue             messaging.Client
}

func NewCreatePaymentUsecase(
	paymentRepository domain_repository.PaymentRepositoryInterface,
	queue messaging.Client,
) domain_payment_usecase.CreatePaymentUsecaseInterface {
	return &CreatePaymentUsecase{
		PaymentRepository: paymentRepository,
		Queue:             queue,
	}
}

func (u *CreatePaymentUsecase) Execute(ctx context.Context, dto dto.CreatePaymentDto) error {
	paymentAlreadyProcessed := u.PaymentRepository.AlreadyAdded(ctx, dto.CorrelationId)

	if paymentAlreadyProcessed {
		return errors.New("payment already processed")
	}

	return nil
}
