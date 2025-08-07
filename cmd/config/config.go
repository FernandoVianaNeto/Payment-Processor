package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ApplicationCfg                    *ApplicationConfig
	NatsCfg                           *NatsConfig
	PaymentProcessorDefaultClientCfg  *PaymentProcessorClientConfig
	PaymentProcessorFallbackClientCfg *PaymentProcessorClientConfig
	MongoCfg                          *MongoConfig
)

const (
	AppName     = "payment-gateway-Backend"
	AppVersion  = "1.0.0"
	Development = "development"
	Staging     = "stage"
	Production  = "production"
)

type ApplicationConfig struct {
	Env         string
	AppVersion  string
	AppPort     int
	Environment string
}

type NatsConfig struct {
	Host                 string
	User                 string
	Password             string
	PaymentRequestsQueue string
}

type MongoConfig struct {
	PaymentCollection string
	Dsn               string
	Database          string
}

type PaymentProcessorClientConfig struct {
	BaseUri string
}

func initialize() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func InitializeConfigs() {
	initialize()
	initializeApplicationConfigs()
	initializeNatsConfigs()
	initializeProcessorsPayment()
	initializeMongoConfig()
}

func getEnv(key string, defaultVal string) string {
	value, exists := os.LookupEnv(key)

	if exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")

	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func initializeApplicationConfigs() {
	if ApplicationCfg == nil {
		ApplicationCfg = &ApplicationConfig{
			Env:        getEnv("APP_ENV", "local"),
			AppVersion: AppVersion,
			AppPort:    getEnvAsInt("APP_PORT", 8080),
		}
	}
}

func initializeNatsConfigs() {
	if NatsCfg == nil {
		NatsCfg = &NatsConfig{
			Host:                 getEnv("NATS_HOST", "nats://nats:4222"),
			User:                 getEnv("NATS_USER", ""),
			Password:             getEnv("NATS_PASSWORD", ""),
			PaymentRequestsQueue: getEnv("NATS_PAYMENT_REQUEST_QUEUE", "payment_requests"),
		}
	}
}

func initializeMongoConfig() {
	if MongoCfg == nil {
		MongoCfg = &MongoConfig{
			PaymentCollection: getEnv("MONGO_PAYMENT_COLLECTION", "payment-gateway"),
			Dsn:               getEnv("MONGO_DSN", "mongodb://mongo-db:27017"),
			Database:          getEnv("MONGO_DATABASE", "payment-gateway"),
		}
	}
}

func initializeProcessorsPayment() {
	if PaymentProcessorDefaultClientCfg == nil {
		PaymentProcessorDefaultClientCfg = &PaymentProcessorClientConfig{
			BaseUri: getEnv("PROCESSOR_DEFAULT_BASE_URI", "http://127.0.0.1:8001"),
		}
	}

	if PaymentProcessorFallbackClientCfg == nil {
		PaymentProcessorFallbackClientCfg = &PaymentProcessorClientConfig{
			BaseUri: getEnv("PROCESSOR_FALLBACK_BASE_URI", "http://127.0.0.1:8002"),
		}
	}
}
