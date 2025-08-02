package domain_payment_usecase

import (
	"context"
)

type PaymentsSummaryUsecaseInterface interface {
	Execute(ctx context.Context) error
}
