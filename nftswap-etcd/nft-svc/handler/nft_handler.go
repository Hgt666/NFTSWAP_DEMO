package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"nftswap-etcd/nft-svc/models"
	"nftswap-etcd/nft-svc/service"
	"nftswap-etcd/pkg/cache"
	"nftswap-etcd/pkg/util"

	"github.com/gin-gonic/gin"
)

func NftListHandler(c *gin.Context) {
	var req models.NftListReq
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, models.Resp{
			Code:    400,
			Message: "参数错误",
		})
		return
	}

	ctx := context.Background()
	cacheKey := util.GenNftListKey(req.Page, req.PageSize)

	// 1. 查本地 ristretto 缓存
	localData, ok := cache.Get(cacheKey)
	if ok {
		c.JSON(http.StatusOK, models.Resp{
			Code:    200,
			Message: "success",
			Data:    localData,
		})
		return
	}

	// 2. 查分布式 Redis
	redisData, err := cache.RedisGet(ctx, cacheKey)
	if err == nil && redisData != "" {
		var data gin.H
		_ = json.Unmarshal([]byte(redisData), &data)
		// 回写本地缓存
		cache.Set(cacheKey, data)
		c.JSON(http.StatusOK, models.Resp{
			Code:    200,
			Message: "success",
			Data:    data,
		})
		return
	}

	// 3. 查数据库
	list, total, err := service.GetNftList(ctx, req.Page, req.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, models.Resp{
			Code:    500,
			Message: "查询失败",
		})
		return
	}

	respData := gin.H{
		"list":  list,
		"total": total,
	}

	// 4. 写入两级缓存
	dataBytes, _ := json.Marshal(respData)
	_ = cache.RedisSet(ctx, cacheKey, string(dataBytes), 60)
	cache.Set(cacheKey, respData)

	c.JSON(http.StatusOK, models.Resp{
		Code:    200,
		Message: "success",
		Data:    respData,
	})
}