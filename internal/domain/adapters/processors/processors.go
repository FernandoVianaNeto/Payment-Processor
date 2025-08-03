package processors

import "context"

type ProcessorClientInput struct {
	CorrelationId string  `json:"correlationId"`
	Amount        float64 `json:"amount"`
	RequestedAt   string  `json:"requestedAt"`
}

type ProcessorsClientInterface interface {
	ExecutePayment(ctx context.Context, input ProcessorClientInput) error
	HealthCheck(ctx context.Context) error
}
