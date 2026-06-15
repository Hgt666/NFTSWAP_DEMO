package utils

import (

	"github.com/dgraph-io/ristretto"
)

// 缓存
var LocalCache *ristretto.Cache

// 初始化缓存
func InitCache() error {
	cache,err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 10000,
		MaxCost: 64<<20,// 最大内存 256MB
		BufferItems: 128,
	})
	if err != nil {
		return err
	}
	LocalCache = cache
	return nil

	
}