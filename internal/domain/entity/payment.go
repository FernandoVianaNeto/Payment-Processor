package entity

type Payment struct {
	CorrelationId        string  `json:"correlationId"`
	RequestedAt          string  `json:"requestedAt"`
	Amount               float64 `json:"amount"`
	TransactionProcessor string  `json:"transactionProcessor"`
}

func NewPayment(
	correlationId string,
	amount float64,
) *Payment {
	entity := &Payment{
		CorrelationId: correlationId,
		Amount:        amount,
	}
	return entity
}
