package api

import (
	"nftswap-etcd/nft-svc/handler"
	"github.com/gin-gonic/gin"
)

func RegisterNftRouter(r *gin.Engine) {
	apiGroup := r.Group("/api/v1")
	{
		apiGroup.GET("/nft/list", handler.NftListHandler)
	}
}