package payment_usecase

import (
	"context"
	"payment-gateway/internal/domain/dto"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_response "payment-gateway/internal/domain/response"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"
)

type GetPaymentsSummaryUsecase struct {
	PaymentRepository domain_repository.PaymentRepositoryInterface
}

func NewGetPaymentsSummaryUsecase(
	paymentRepository domain_repository.PaymentRepositoryInterface,
) domain_payment_usecase.GetPaymentsSummaryUsecaseInterface {
	return &GetPaymentsSummaryUsecase{
		PaymentRepository: paymentRepository,
	}
}

func (u *GetPaymentsSummaryUsecase) Execute(ctx context.Context, query dto.GetPaymentsSummaryDto) (*domain_response.PaymentSummaryResponse, error) {
	summaryAggregated, err := u.PaymentRepository.Summary(ctx, query)

	if err != nil {
		return nil, err
	}

	return summaryAggregated, err
}
