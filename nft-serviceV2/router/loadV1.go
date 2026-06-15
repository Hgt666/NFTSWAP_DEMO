package router

import (
	"nft-service/handler"

	"github.com/gin-gonic/gin"
)

// 分版路由
func LoadV1(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.GET("/ping",func(c *gin.Context) {
			c.JSON(200,gin.H{"message":"pong"})
		})

		// 查找所有挂单的 nft
		api.GET("/nft/list",handler.NFTlistHandlerLocalCache)
	}

}
