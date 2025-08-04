package redis_payment_repository

import (
	"context"
	"payment-gateway/internal/domain/dto"
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

func (f *PaymentsRepository) Create(ctx context.Context, input dto.CreatePaymentDto) error {
	response := f.client.Set(ctx, input.CorrelationId, input.Amount, 0)

	return response.Err()
}

func (f *PaymentsRepository) Summary(ctx context.Context) error {
	return nil
}

func (f *PaymentsRepository) AlreadyAdded(ctx context.Context, correlationId string) bool {
	payment := f.client.Get(ctx, correlationId)

	if payment.Val() != "" {
		return true
	}

	return false
}
