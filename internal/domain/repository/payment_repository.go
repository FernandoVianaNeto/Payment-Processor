package domain_repository

import (
	"context"
)

//go:generate mockgen -source $GOFILE -package $GOPACKAGE -destination $ROOT_DIR/test/mocks/$GOPACKAGE/mock_$GOFILE

type PaymentRepositoryInterface interface {
	Create(ctx context.Context) error
	Summary(ctx context.Context) error
}
