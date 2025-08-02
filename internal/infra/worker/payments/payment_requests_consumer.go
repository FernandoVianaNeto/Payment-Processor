package payment_request_queue_consumer

import (
	"context"
	"log"
	configs "payment-gateway/cmd/config"
	"payment-gateway/internal/domain/adapters/queue"
)

func StartPaymentRequestsConsumer(ctx context.Context, client queue.Client, consumerName string) error {
	err := client.Subscribe(configs.NatsCfg.PaymentRequestsQueue, func(msg []byte) {
		log.Printf("[consumer1] Mensagem recieved: %s\n", string(msg))
		// Processar mensagem aqui
	})
	if err != nil {
		return err
	}

	log.Println("✅ Worker de workout iniciado e aguardando mensagens...")

	<-ctx.Done()
	log.Println("🛑 Worker workout encerrado via contexto.")
	return nil
}
