package repository

import (
	"github.com/redis/go-redis/v9"
)

func RedisInit() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return client
}
