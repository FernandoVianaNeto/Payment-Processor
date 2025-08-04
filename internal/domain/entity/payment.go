package entity

type Payment struct {
	CorrelationId string  `json:"correlationId"`
	RequestedAt   string  `json:"requestedAt"`
	Amount        float64 `json:"amount"`
}

func NewPayment(
	correlationId string,
	amount float64,
	RequestedAt string,
) *Payment {
	entity := &Payment{
		CorrelationId: correlationId,
		Amount:        amount,
		RequestedAt:   RequestedAt,
	}
	return entity
}
