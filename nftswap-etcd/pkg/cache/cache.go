package cache

import (
	"github.com/dgraph-io/ristretto"
	"github.com/go-redis/redis/v8"
	"context"
)

// LocalCacheConfig 本地缓存配置
type LocalCacheConfig struct {
	NumCounters int64 `yaml:"num_counters"`
	MaxCost     int64 `yaml:"max_cost"`
	BufferItems int64 `yaml:"buffer_items"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Addr string `yaml:"addr"`
}

var (
	LocalCache *ristretto.Cache
	RedisCli   *redis.Client
)

// InitLocalCache 初始化 ristretto
func InitLocalCache(cfg LocalCacheConfig) error {
	rcfg := &ristretto.Config{
		NumCounters: cfg.NumCounters,
		MaxCost:     cfg.MaxCost,
		BufferItems: cfg.BufferItems,
	}
	cache, err := ristretto.NewCache(rcfg)
	if err != nil {
		return err
	}
	LocalCache = cache
	return nil
}

// InitRedis 初始化 Redis
func InitRedis(cfg RedisConfig) {
	RedisCli = redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
	})
}

// Get 读取本地缓存
func Get(key string) (interface{}, bool) {
	return LocalCache.Get(key)
}

// Set 写入本地缓存
func Set(key string, val interface{}) bool {
	return LocalCache.Set(key, val, 1)
}

// Del 删除本地缓存
func Del(key string) {
	LocalCache.Del(key)
}

// RedisGet Redis读取
func RedisGet(ctx context.Context, key string) (string, error) {
	return RedisCli.Get(ctx, key).Result()
}

// RedisSet Redis写入
func RedisSet(ctx context.Context, key string, val string, expire int) error {
	return RedisCli.Set(ctx, key, val, 0).Err()
}