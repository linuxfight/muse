package db

import (
	"context"
	"muse/internal/services/logger"
	"os"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	db *redis.Client
}

func New(ctx context.Context) *Service {
	addr := os.Getenv("REDIS_ADDR")

	if addr == "" {
		logger.Log.Fatal("REDIS_ADDR environment variable not set")
		return nil
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	if err := rdb.Ping(ctx).Err(); err != nil {
		logger.Log.Fatal("Error connecting to redis")
	}

	return &Service{
		db: rdb,
	}
}
