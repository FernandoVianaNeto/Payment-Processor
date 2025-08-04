package requests

type CreatePaymentRequest struct {
	CorrelationId string  `json:"correlationId"`
	Amount        float64 `json:"amount"`
}

type GetSummaryRequet struct {
	From string `form:"from"`
	To   string `form:"to"`
}
