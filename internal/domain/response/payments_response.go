package domain_response

type Summary struct {
	TotalRequests int     `json:"totalRequests"`
	TotalAmount   float64 `json:"totalAmount"`
}
type PaymentSummaryResponse struct {
	Default  Summary `json:"default"`
	Fallback Summary `json:"fallback"`
}

type HealthCheckResponse struct {
	Failing         bool `json:"failing"`
	MinResponseTime int  `json:"minResponseTime"`
}
