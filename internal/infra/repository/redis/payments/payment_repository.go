package redis_payment_repository

import (
	"context"
	"payment-gateway/internal/domain/dto"
	"payment-gateway/internal/domain/entity"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_response "payment-gateway/internal/domain/response"

	jsoniter "github.com/json-iterator/go"
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
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	entity := entity.NewPayment(input.CorrelationId, input.Amount, input.RequestedAt)

	entityEncoded, _ := json.Marshal(entity)

	response := f.client.Set(ctx, input.CorrelationId, entityEncoded, 0)

	return response.Err()
}

func (f *PaymentsRepository) AlreadyAdded(ctx context.Context, correlationId string) bool {
	payment := f.client.Get(ctx, correlationId)

	if payment.Val() != "" {
		return true
	}

	return false
}

func (f *PaymentsRepository) Summary(ctx context.Context, input dto.GetPaymentsSummaryDto) (*domain_response.PaymentSummaryResponse, error) {
	return nil, nil
}
