package domain_payment_usecase

import (
	"context"
)

type ProcessPaymentRequestUsecaseInterface interface {
	Execute(ctx context.Context, message []byte) error
}
