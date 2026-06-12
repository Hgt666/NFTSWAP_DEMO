package db

import (
	"context"

	"github.com/go-redis/redis/v8"
	"nft-service/config"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: config.RedisAddr,
	})
	return RedisClient.Ping(Ctx).Err()
}