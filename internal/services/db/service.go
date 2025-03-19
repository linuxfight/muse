package db

import (
	"context"
	"muse/internal/services/logger"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	db *redis.Client
}

func New(ctx context.Context, addr string) *Service {
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
