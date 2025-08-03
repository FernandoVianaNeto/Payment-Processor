package processors

import (
	"context"
	"encoding/json"
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
	// paymentAlreadyProcessed := u.PaymentRepository.AlreadyAdded(ctx, dto.CorrelationId)

	// if paymentAlreadyProcessed {
	// 	return errors.New("payment already processed")
	// }

	// requestedAt := time.Now().UTC().String()

	// message := entity.NewPayment(dto.CorrelationId, dto.Amount, requestedAt)

	// messageByte, err := json.Marshal(message)

	// if err != nil {
	// 	return err
	// }

	// err = u.Queue.Publish(configs.NatsCfg.PaymentRequestsQueue, messageByte)

	// return err

	return nil
}

func (u *ProcessorDefaultClient) IsLive(ctx context.Context) (*domain_response.HealthCheckResponse, error) {
	var response domain_response.HealthCheckResponse

	responseByte, statusCode, err := u.Client.Get(ctx, "/payment/service-health", url.Values{})

	if err != nil || statusCode != http.StatusOK {
		return nil, err
	}

	err = json.Unmarshal(responseByte, &response)

	if err != nil {
		return nil, err
	}

	return &response, err
}
