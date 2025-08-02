package requests

type GetSummary struct {
}

type CreatePaymentRequest struct {
	CorrelationId string  `json:"correlationId"`
	Amount        float64 `json:"amount"`
}
