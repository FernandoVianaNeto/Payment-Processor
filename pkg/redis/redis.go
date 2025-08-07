package redis_client

import (
	"context"
	"fmt"
	"log"
	configs "payment-gateway/cmd/config"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	RDB *redis.Client
)

func InitRedis() *redis.Client {
	configs.InitializeConfigs()

	addr := fmt.Sprintf("%s:%d", configs.RedisCfg.Host, 6379)

	RDB = redis.NewClient(&redis.Options{
		Addr: addr,
		// Password:     configs.RedisCfg.Password,
		// DB:           configs.RedisCfg.Db,
		// MinIdleConns: configs.RedisCfg.MinIddleConns,
		// PoolSize:     configs.RedisCfg.PoolSize,
	})

	if err := RDB.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	log.Println("✅ Connected to Redis at", addr)
	return RDB
}
