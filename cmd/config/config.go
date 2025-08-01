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
)

const (
	AppName     = "forfit-Backend"
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
	Host              string
	User              string
	Password          string
	WorkoutTopic      string
	UserTopic         string
	GamificationTopic string
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
			Host:              getEnv("NATS_HOST", "nats://localhost:4222"),
			User:              getEnv("NATS_USER", "root"),
			Password:          getEnv("NATS_PASSWORD", "password"),
			UserTopic:         getEnv("NATS_USER_TOPIC", "user.events"),
			WorkoutTopic:      getEnv("NATS_WORKOUT_TOPIC", "workout.events"),
			GamificationTopic: getEnv("NATS_GAMIFICATION_TOPIC", "gamification.events"),
		}
	}
}
