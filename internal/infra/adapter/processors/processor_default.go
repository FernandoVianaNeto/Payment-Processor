package processors

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"payment-gateway/internal/domain/adapters/processors"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_response "payment-gateway/internal/domain/response"
	http_client "payment-gateway/pkg/client/http"
)

type ProcessorDefaultClient struct {
	PaymentRepository domain_repository.PaymentRepositoryInterface
	Client            *http_client.Client
}

func NewProcessorDefaultClient(
	client *http_client.Client,
	paymentRepository domain_repository.PaymentRepositoryInterface,
) processors.ProcessorsClientInterface {
	return &ProcessorDefaultClient{
		Client:            client,
		PaymentRepository: paymentRepository,
	}
}

func (u *ProcessorDefaultClient) ExecutePayment(ctx context.Context, input processors.ProcessorClientInput) error {
	body, err := json.Marshal(input)

	if err != nil {
		return err
	}

	_, statusCode, err := u.Client.Post(ctx, "/payments", url.Values{}, body)

	if statusCode != http.StatusOK || err != nil {
		return errors.New("error processing payment")
	}

	return err
}

func (u *ProcessorDefaultClient) IsLive(ctx context.Context) (*domain_response.HealthCheckResponse, error) {
	var response domain_response.HealthCheckResponse

	responseByte, statusCode, err := u.Client.Get(ctx, "/payments/service-health", url.Values{})

	if err != nil || statusCode != http.StatusOK {
		return nil, err
	}

	err = json.Unmarshal(responseByte, &response)

	if err != nil {
		return nil, err
	}

	return &response, err
}
