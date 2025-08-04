package domain_payment_usecase

import (
	"context"
	"payment-gateway/internal/domain/dto"
	domain_response "payment-gateway/internal/domain/response"
)

type GetPaymentsSummaryUsecaseInterface interface {
	Execute(ctx context.Context, query dto.GetPaymentsSummaryDto) (*domain_response.PaymentSummaryResponse, error)
}
