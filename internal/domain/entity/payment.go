package entity

type Payment struct {
	CorrelationId string  `json:"correlation_id"`
	CreatedAt     string  `json:"created_at"`
	Amount        float64 `json:"amount"`
}

func NewPayment(
	correlationId string,
	amount float64,
	createdAt string,
) *Payment {
	entity := &Payment{
		CorrelationId: correlationId,
		Amount:        amount,
		CreatedAt:     createdAt,
	}
	return entity
}
