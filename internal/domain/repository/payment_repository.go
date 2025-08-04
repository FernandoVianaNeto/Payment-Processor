package domain_repository

import (
	"context"
	"payment-gateway/internal/domain/dto"
)

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

type PaymentRepositoryInterface interface {
	Create(ctx context.Context, input dto.CreatePaymentDto) error
	AlreadyAdded(ctx context.Context, correlationId string) bool
	Summary(ctx context.Context) error
}
