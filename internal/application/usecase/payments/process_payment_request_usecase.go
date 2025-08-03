package payment_usecase

import (
	"context"
	"payment-gateway/internal/domain/adapters/processors"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"
)

type ProcessPaymentRequestUsecase struct {
	PaymentRepository             domain_repository.PaymentRepositoryInterface
	ProcessPaymentDefaultAdapter  processors.ProcessorsClientInterface
	ProcessPaymentFallbackAdapter processors.ProcessorsClientInterface
}

func NewProcessPaymentRequestUsecase(
	paymentRepository domain_repository.PaymentRepositoryInterface,
	processPaymentDefaultAdapter processors.ProcessorsClientInterface,
	processPaymentFallbackAdapter processors.ProcessorsClientInterface,
) domain_payment_usecase.ProcessPaymentRequestUsecaseInterface {
	return &ProcessPaymentRequestUsecase{
		PaymentRepository:             paymentRepository,
		ProcessPaymentDefaultAdapter:  processPaymentDefaultAdapter,
		ProcessPaymentFallbackAdapter: processPaymentFallbackAdapter,
	}
}

func (u *ProcessPaymentRequestUsecase) Execute(ctx context.Context, data []byte) error {
	return nil
}
