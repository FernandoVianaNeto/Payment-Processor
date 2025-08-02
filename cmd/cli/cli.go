package cli

import (
	"context"
	"fmt"
	"os"
	configs "payment-gateway/cmd/config"
	payment_request_queue_consumer "payment-gateway/internal/infra/worker/payments"
	natsclient "payment-gateway/pkg/nats"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(paymentRequestsConsumerCmd)

}

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "root - main command application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error executing command: %v\n", err)
		os.Exit(1)
	}
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Initialize Http Server",
	Run: func(cmd *cobra.Command, args []string) {
		port := configs.ApplicationCfg.AppPort
		if port == 0 {
			os.Exit(1)
		}

		srv := NewApplication()
		if err := srv.Start(fmt.Sprintf(":%d", port)); err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing application: %v\n", err)
			os.Exit(1)
		}
	},
}

var paymentRequestsConsumerCmd = &cobra.Command{
	Use:   "payment-requests-consumer",
	Short: "Initialize Payment Requests Consumer",
	Run: func(cmd *cobra.Command, args []string) {
		configs.InitializeConfigs()

		ctx := context.Background()

		client := natsclient.New(configs.NatsCfg.Host)

		client.Connect()

		err := payment_request_queue_consumer.StartPaymentRequestsConsumer(ctx, client, "payment_requests_consumer")

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to start payment_requests_consumer: %v\n", err)
			os.Exit(1)
			return
		}

		fmt.Println("payment_requests_consumer started successfully")
	},
}
