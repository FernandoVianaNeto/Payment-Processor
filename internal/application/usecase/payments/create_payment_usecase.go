package payment_usecase

import (
	"context"
	"payment-gateway/internal/domain/dto"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"
)

type CreatePaymentUsecase struct {
	PaymentRepository domain_repository.PaymentRepositoryInterface
}

func NewCreatePaymentUsecase(
	paymentRepository domain_repository.PaymentRepositoryInterface,
) domain_payment_usecase.CreatePaymentUsecaseInterface {
	return &CreatePaymentUsecase{
		PaymentRepository: paymentRepository,
	}
}

func (u *CreatePaymentUsecase) Execute(ctx context.Context, dto dto.CreatePaymentDto) error {
	return nil
}
