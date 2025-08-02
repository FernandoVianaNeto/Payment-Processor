package redis_payment_repository

import (
	"context"
	domain_repository "payment-gateway/internal/domain/repository"

	"github.com/redis/go-redis/v9"
)

type PaymentsRepository struct {
	client *redis.Client
}

func NewPaymentsRepository(client *redis.Client) domain_repository.PaymentRepositoryInterface {
	return &PaymentsRepository{
		client: client,
	}
}

func (f *PaymentsRepository) Create(ctx context.Context) error {
	return nil
}

func (f *PaymentsRepository) Summary(ctx context.Context) error {
	return nil
}

func (f *PaymentsRepository) AlreadyAdded(ctx context.Context, correlationId string) bool {
	payment := f.client.Get(ctx, correlationId)

	if payment != nil {
		return true
	}

	return false
}
