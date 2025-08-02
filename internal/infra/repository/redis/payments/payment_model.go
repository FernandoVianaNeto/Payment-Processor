package redis_payment_repository

type PaymentModel struct {
	CorrelationId string  `bson:"correlation_id" json:"correlation_id"`
	CreatedAt     string  `json:"created_at"`
	Amount        float64 `json:"amount"`
}
