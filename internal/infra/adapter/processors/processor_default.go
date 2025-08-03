package processors

import (
	"context"
	"payment-gateway/internal/domain/adapters/processors"
	domain_repository "payment-gateway/internal/domain/repository"
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

func (u *ProcessorDefaultClient) HealthCheck(ctx context.Context) error {
	return nil
}
