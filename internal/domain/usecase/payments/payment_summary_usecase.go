package domain_gamification_usecase

import (
	"context"
)

type PaymentsSummaryUsecaseInterface interface {
	Execute(ctx context.Context) error
}
