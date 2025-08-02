package domain_payment_usecase

import (
	"context"
	"payment-gateway/internal/domain/dto"
)

type CreatePaymentUsecaseInterface interface {
	Execute(ctx context.Context, dto dto.CreatePaymentDto) error
}
