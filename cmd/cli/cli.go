package cli

import (
	"fmt"
	"os"
	configs "payment-gateway/cmd/config"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(httpCmd)
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

// var workoutConsumerCmd = &cobra.Command{
// 	Use:   "workout-consumer",
// 	Short: "Initialize Workout Consumer",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		configs.InitializeConfigs()

// 		ctx := context.Background()

// 		client := natsclient.New(configs.NatsCfg.Host)

// 		client.Connect()

// 		err := worker_workout.StartWorkoutConsumer(ctx, client, "workout_consumer")

// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "Failed to start workout consumer: %v\n", err)
// 			os.Exit(1)
// 			return
// 		}

// 		fmt.Println("Workout consumer started successfully")
// 	},
// }
