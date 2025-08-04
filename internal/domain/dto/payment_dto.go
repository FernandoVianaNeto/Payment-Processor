package dto

type CreatePaymentDto struct {
	CorrelationId        string  `json:"correlationId"`
	Amount               float64 `json:"amount"`
	RequestedAt          string  `json:"requestedAt"`
	TransactionProcessor string  `json:"transactionProcessor"`
}

type ProcessPaymentRequestDto struct {
	CorrelationId string  `json:"correlationId"`
	Amount        float64 `json:"amount"`
	RequestedAt   string  `json:"requestedAt"`
	Retries       int     `json:"retries"`
}

type GetPaymentsSummaryDto struct {
	From string `json:"from"`
	To   string `json:"to"`
}
