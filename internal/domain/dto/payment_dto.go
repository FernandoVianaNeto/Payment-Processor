package dto

type CreatePaymentDto struct {
	CorrelationId string  `json:"correlation_id"`
	Amount        float64 `json:"amount"`
}
