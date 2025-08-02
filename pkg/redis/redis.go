package redis_client

import (
	"context"
	"log"
	configs "payment-gateway/cmd/config"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	RDB *redis.Client
)

func InitRedis() *redis.Client {
	host := configs.RedisCfg.Host

	RDB = redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     configs.RedisCfg.Password,
		DB:           configs.RedisCfg.Db,
		MinIdleConns: configs.RedisCfg.MinIddleConns,
		PoolSize:     configs.RedisCfg.PoolSize,
	})

	if err := RDB.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	log.Println("✅ Connected to Redis at", host)
	return RDB
}
