package main

import (
	"payment-gateway/cmd/cli"
	configs "payment-gateway/cmd/config"
)

func main() {
	configs.InitializeConfigs()

	cli.Execute()
}
