package worker_workout

import (
	"context"
	"log"
	configs "payment-gateway/cmd/config"
	"payment-gateway/internal/domain/adapters/messaging"
)

func StartWorkoutConsumer(ctx context.Context, client messaging.Client, consumerName string) error {
	err := client.Subscribe(configs.NatsCfg.WorkoutTopic, func(msg []byte) {
		log.Printf("[consumer1] Mensagem recebida: %s\n", string(msg))
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
