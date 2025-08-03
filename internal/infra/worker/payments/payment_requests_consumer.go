package payment_request_queue_consumer

import (
	"context"
	"log"
	configs "payment-gateway/cmd/config"
	"payment-gateway/internal/domain/adapters/processors"
	"payment-gateway/internal/domain/adapters/queue"
	domain_repository "payment-gateway/internal/domain/repository"
	domain_payment_usecase "payment-gateway/internal/domain/usecase/payments"
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
	err := client.Subscribe(configs.NatsCfg.PaymentRequestsQueue, func(msg []byte) {
		log.Printf("[consumer1] Mensagem recieved: %s\n", string(msg))
		// Processar mensagem aqui
	})
	if err != nil {
		return err
	}

	log.Println("Payment Request Consumer sucessfully started")

	<-ctx.Done()
	log.Println("🛑 Worker workout encerrado via contexto.")
	return nil
}
