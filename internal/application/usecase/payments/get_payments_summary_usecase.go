package payment_usecase

import (
	"context"
	"payment-gateway/internal/domain/adapters/queue"
	"payment-gateway/internal/domain/dto"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_response "payment-gateway/internal/domain/response"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"
)

type GetPaymentsSummaryUsecase struct {
	PaymentRepository domain_repository.PaymentRepositoryInterface
	Queue             queue.Client
}

func NewGetPaymentsSummaryUsecase(
	paymentRepository domain_repository.PaymentRepositoryInterface,
	queue queue.Client,
) domain_payment_usecase.GetPaymentsSummaryUsecaseInterface {
	return &GetPaymentsSummaryUsecase{
		PaymentRepository: paymentRepository,
		Queue:             queue,
	}
}

func (u *GetPaymentsSummaryUsecase) Execute(ctx context.Context, query dto.GetPaymentsSummaryDto) (*domain_response.PaymentSummaryResponse, error) {
	summaryAggregated, err := u.PaymentRepository.Summary(ctx, query)

	if err != nil {
		return nil, err
	}

	return summaryAggregated, err
}
