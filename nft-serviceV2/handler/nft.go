package handler

import (
	"net/http"
	"nft-service/models"
	"nft-service/service"

	"github.com/gin-gonic/gin"
)

// 获取所有挂单的nft
func NFTlistHandler(c *gin.Context)  {
	var req models.NftListReq
	// 绑定并校验参数
	if err := c.ShouldBindQuery(&req);err != nil{
		c.JSON(http.StatusOK,models.Resp{
			Code: 400,
			Msg: "参数错误",
			Data: nil,
		})
		return
	}

	// 调用业务层
	list,total,err := service.GetNftList(req)
	if err != nil {
		c.JSON(http.StatusOK,models.Resp{
			Code: 500,
			Msg: "查询失败",
			Data: nil,
		})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK,models.Resp{
		Code: 200,
		Msg: "查询成功",
		Data: gin.H{
			"list": list,
			"total": total,
		},
	})
	
}