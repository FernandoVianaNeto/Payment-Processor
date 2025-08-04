package payment_request_queue_consumer

import (
	"context"
	"log"
	configs "payment-gateway/cmd/config"
	"payment-gateway/internal/domain/adapters/processors"
	"payment-gateway/internal/domain/adapters/queue"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"

	"golang.org/x/sync/semaphore"
)

type Usecases struct {
	ProcessPaymentRequestUsecase domain_payment_usecase.ProcessPaymentRequestUsecaseInterface
}

type ConsumerInfra struct {
	ProcessorPaymentDefault  processors.ProcessorsClientInterface
	ProcessorPaymentFallback processors.ProcessorsClientInterface
	PaymentRepository        domain_repository.PaymentRepositoryInterface
	Usecases                 Usecases
}

func StartPaymentRequestsConsumer(ctx context.Context, client queue.Client, consumerName string, consumerInfra ConsumerInfra) error {
	semaphore := semaphore.NewWeighted(int64(10))

	err := client.Subscribe(configs.NatsCfg.PaymentRequestsQueue, func(msg []byte) {
		log.Printf("[payment-gateway-processor] Mensagem received: %s\n", string(msg))

		if err := semaphore.Acquire(ctx, 1); err != nil {
			log.Println("Could not acquire semaphore: ", err)
		}

		go func(data []byte) {
			defer semaphore.Release(1)

			err := consumerInfra.Usecases.ProcessPaymentRequestUsecase.Execute(ctx, data)

			if err != nil {
				log.Println("error on executing process payment request: ", err)
				return
			}
		}(msg)
	})

	if err != nil {
		return err
	}

	log.Println("Payment Request Consumer sucessfully started")

	<-ctx.Done()
	log.Println("Payment Request Consumer canceled due Context")
	return nil
}
