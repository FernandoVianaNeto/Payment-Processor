package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ApplicationCfg *ApplicationConfig
	NatsCfg        *NatsConfig
	RedisCfg       *RedisConfig
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

type RedisConfig struct {
	Host          string
	Port          int
	Password      string
	Db            int
	MinIddleConns int
	PoolSize      int
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
	initializeRedisConfig()
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
			AppPort:    getEnvAsInt("APP_PORT", 3001),
		}
	}
}

func initializeNatsConfigs() {
	if NatsCfg == nil {
		NatsCfg = &NatsConfig{
			Host:                 getEnv("NATS_HOST", "nats://localhost:4222"),
			User:                 getEnv("NATS_USER", "root"),
			Password:             getEnv("NATS_PASSWORD", "password"),
			PaymentRequestsQueue: getEnv("NATS_PAYMENT_REQUEST_QUEUE", "payment_requests"),
		}
	}
}

func initializeRedisConfig() {
	if RedisCfg == nil {
		RedisCfg = &RedisConfig{
			Host:          getEnv("REDIS_HOST", "redis"),
			Port:          getEnvAsInt("REDIS_PORT", 6379),
			Password:      getEnv("REDIS_PASSWORD", "password"),
			Db:            getEnvAsInt("REDIS_DB", 0),
			MinIddleConns: getEnvAsInt("REDIS_MIN_IDDLE_CONNS", 1),
			PoolSize:      getEnvAsInt("REDIS_POOL_SIZE", 5),
		}
	}
}
