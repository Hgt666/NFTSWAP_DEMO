package common

import (
	"nftswap-etcd/pkg/cache"
	"nftswap-etcd/nft-svc/db"
)

// 引用外部配置结构体
type (
	MysqlConfig     = db.MysqlConfig
	LocalCacheConfig = cache.LocalCacheConfig
	RedisConfig     = cache.RedisConfig
)

// InitAll 统一初始化 MySQL、Redis、本地缓存
func InitAll(mysqlCfg MysqlConfig, redisCfg RedisConfig, localCacheCfg LocalCacheConfig) error {
	// 初始化MySQL
	if err := db.InitMySQL(mysqlCfg); err != nil {
		return err
	}
	// 初始化Redis
	cache.InitRedis(redisCfg)
	// 初始化本地缓存
	if err := cache.InitLocalCache(localCacheCfg); err != nil {
		return err
	}
	return nil
}