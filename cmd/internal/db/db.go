package db

import (
	"context"

	"github.com/Rajit-Dutta/go-redis-url-shortener/cmd/internal/config"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func CreateClient(dbNum int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.MustLoad().DatabaseURL,
		Password: "",
		DB:       dbNum,
	})

	return rdb
}
