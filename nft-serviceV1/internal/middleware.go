package internal

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"nft-service/config"
	// "nft-service/db"
)

// API 鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-API-Token")
		if token != config.APIToken {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "非法访问"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// 简易限流：单IP 1秒最多5次请求
var ipLimit = make(map[string]int)
var mu sync.Mutex

func RateLimitMiddleware() gin.HandlerFunc {
	// 定时清空计数器
	go func() {
		for range time.Tick(1 * time.Second) {
			mu.Lock()
			ipLimit = make(map[string]int)
			mu.Unlock()
		}
	}()

	return func(c *gin.Context) {
		ip := c.ClientIP()
		mu.Lock()
		cnt := ipLimit[ip]
		if cnt >= 5 {
			mu.Unlock()
			c.JSON(http.StatusTooManyRequests, gin.H{"code": 429, "msg": "请求过于频繁"})
			c.Abort()
			return
		}
		ipLimit[ip] = cnt + 1
		mu.Unlock()
		c.Next()
	}
}