package dto

type CreatePaymentDto struct {
	CorrelationId string  `json:"correlation_id"`
	Amount        float64 `json:"amount"`
}

type ProcessPaymentRequestDto struct {
	CorrelationId string  `json:"correlation_id"`
	Amount        float64 `json:"amount"`
	RequestedAt   string  `json:"requestedAt"`
	Retries       int     `json:"retries"`
}
