package database

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Password: "",
		DB:       dbNo,
	})
	return rdb
}
