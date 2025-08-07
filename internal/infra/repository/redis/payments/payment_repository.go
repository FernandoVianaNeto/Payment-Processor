package redis_payment_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"payment-gateway/internal/domain/dto"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_response "payment-gateway/internal/domain/response"

	"github.com/redis/go-redis/v9"
)

type PaymentsRepository struct {
	client *redis.Client
}

type PaymentModel struct {
	CorrelationId        string  `json:"correlationId"`
	RequestedAt          string  `json:"requestedAt"`
	Amount               float64 `json:"amount"`
	TransactionProcessor string  `json:"transactionProcessor"`
}

func NewPaymentsRepository(client *redis.Client) domain_repository.PaymentRepositoryInterface {
	return &PaymentsRepository{
		client: client,
	}
}

func (r *PaymentsRepository) Create(ctx context.Context, input dto.CreatePaymentDto) error {
	key := fmt.Sprintf("payment:%s", input.CorrelationId)

	data := PaymentModel{
		CorrelationId:        input.CorrelationId,
		RequestedAt:          input.RequestedAt,
		Amount:               input.Amount,
		TransactionProcessor: input.TransactionProcessor,
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := r.client.Set(ctx, key, bytes, 0).Err(); err != nil {
		return err
	}

	// Simula agregação: atualiza contadores separados
	r.client.IncrByFloat(ctx, fmt.Sprintf("summary:%s:amount", input.TransactionProcessor), input.Amount)
	r.client.Incr(ctx, fmt.Sprintf("summary:%s:requests", input.TransactionProcessor))

	return nil
}

func (r *PaymentsRepository) Summary(ctx context.Context, input dto.GetPaymentsSummaryDto) (*domain_response.PaymentSummaryResponse, error) {
	defaultAmount, _ := r.client.Get(ctx, "summary:default:amount").Float64()
	defaultRequests, _ := r.client.Get(ctx, "summary:default:requests").Int64()

	fallbackAmount, _ := r.client.Get(ctx, "summary:fallback:amount").Float64()
	fallbackRequests, _ := r.client.Get(ctx, "summary:fallback:requests").Int64()

	return &domain_response.PaymentSummaryResponse{
		Default: domain_response.Summary{
			TotalAmount:   roundToFixed(defaultAmount, 1),
			TotalRequests: int(defaultRequests),
		},
		Fallback: domain_response.Summary{
			TotalAmount:   roundToFixed(fallbackAmount, 1),
			TotalRequests: int(fallbackRequests),
		},
	}, nil
}

func roundToFixed(val float64, precision int) float64 {
	factor := math.Pow(10, float64(precision))
	return math.Round(val*factor) / factor
}
