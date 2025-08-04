package mongo_infra

type PaymentModel struct {
	CorrelationId        string  `json:"correlationId" bson:"correlationId"`
	RequestedAt          string  `json:"requestedAt" bson:"requestedAt"`
	Amount               float64 `json:"amount" bson:"amount"`
	TransactionProcessor string  `json:"transactionProcessor" bson:"transactionProcessor"`
}
