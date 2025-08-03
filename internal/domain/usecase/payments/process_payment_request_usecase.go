package domain_payment_usecase

import (
	"context"
	"payment-gateway/internal/domain/dto"
)

type ProcessPaymentRequestUsecaseInterface interface {
	Execute(ctx context.Context, dto dto.ProcessPaymentRequestDto) error
}
