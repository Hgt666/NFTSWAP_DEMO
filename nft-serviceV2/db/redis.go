package db

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

// RedisClient 是全局的 Redis 客户端实例
var Rdb *redis.Client

// 初始化Redis连接
func InitRedis(Ctx context.Context) error {
	// 连接Redis服务器
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 如果有密码
		DB:       0,  // 默认数据库
	})
	Rdb = redisClient

	// 检查连接是否成功
	_, err := redisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
		return err
	}
	// log.Printf("Connected to Redis: %s", msg)

	return nil
}
