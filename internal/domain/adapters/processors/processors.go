package processors

import (
	"context"
	domain_response "payment-gateway/internal/domain/response"
)

type ProcessorClientInput struct {
	CorrelationId string  `json:"correlationId"`
	Amount        float64 `json:"amount"`
	RequestedAt   string  `json:"requestedAt"`
}

type ProcessorsClientInterface interface {
	ExecutePayment(ctx context.Context, input ProcessorClientInput) error
	IsLive(ctx context.Context) (*domain_response.HealthCheckResponse, error)
}
