package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"nft-service/db"
	"nft-service/models"
	"nft-service/service"
	// "nft-service/utils"
	"sync"
	"time"

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





var (
	localCache=sync.Map{}
	cacheTTL = 30*time.Second
)


// 加缓存的handler
func NFTlistHandlerCache(c *gin.Context) {
	var req models.NftListReq
	if err := c.ShouldBindQuery(&req);err != nil {
		c.JSON(200,models.Resp{
			Code: 200,
			Msg: "参数错误",
			Data: nil,
		})
		return
	}

	// 分页维度缓存key
	cackeKey := fmt.Sprintf("nft:list:p%d_s%d",req.Page,req.PageSize)

	// 0、先查本地内存缓存
	if val,ok :=localCache.Load(cackeKey);ok {
		c.JSON(200,val.(models.Resp))
		return
	}

	// 1、查redis缓存
	val,err := db.Rdb.Get(c,cackeKey).Result()
	if err == nil {
		var data gin.H
		_ = json.Unmarshal([]byte(val),&data)
		c.JSON(200,models.Resp{
			Code: 200,
			Msg: "success",
			Data: data,
		})
		return
	}


	// 2、缓存未命中，查 DB
	list,total,err :=service.GetNftList(req)
	if err != nil {
		c.JSON(200,models.Resp{
			Code: 500,
			Msg: "查询失败",
			Data: nil,
		})
		return
	}

	respData := gin.H{
		"list":list,
		"total":total,
	}
	respBytes,_ :=json.Marshal(respData)


	// 3、回写缓存，过期时间30-60s，根据业务调整
	// 回写本地缓存+定时过期
	resp := models.Resp{
		Code: 200,
		Msg: "success",
		Data: respData,
	}
	localCache.Store(cackeKey,resp)
	// 异步定时清理（简单实现）
	go func(k string){
		time.Sleep(cacheTTL)
		localCache.Delete(k)
	}(cackeKey)
	// 回写redis缓存
	expire := cacheTTL+ time.Duration(rand.Intn(10) ) * time.Second
	db.Rdb.Set(c,cackeKey,respBytes,expire)
	c.JSON(200,models.Resp{
		Code: 200,
		Msg: "success",
		Data: respData,
	})
}


// 加本地缓存的handler
func NFTlistHandlerLocalCache(c *gin.Context) {
	var req models.NftListReq
	if err := c.ShouldBindQuery(&req);err != nil {
		c.JSON(200,models.Resp{
			Code: 200,
			Msg: "参数错误",
			Data: nil,
		})
		return
	}

	// 分页维度缓存key
	cackeKey := fmt.Sprintf("nft:list:p%d_s%d",req.Page,req.PageSize)

	// 0、先查本地内存缓存
	if val,ok :=localCache.Load(cackeKey);ok {
		c.JSON(200,val.(models.Resp))
		return
	}
	// if val,ok := utils.LocalCache.Get(cackeKey);ok {
	// 	c.JSON(200,val.(models.Resp))
	// 	return
	// }

	// 1、查redis缓存
	val,err := db.Rdb.Get(c,cackeKey).Result()
	if err == nil {
		var data gin.H
		_ = json.Unmarshal([]byte(val),&data)
		c.JSON(200,models.Resp{
			Code: 200,
			Msg: "success",
			Data: data,
		})
		return
	}


	// 2、缓存未命中，查 DB
	list,total,err :=service.GetNftList(req)
	if err != nil {
		c.JSON(200,models.Resp{
			Code: 500,
			Msg: "查询失败",
			Data: nil,
		})
		return
	}

	respData := gin.H{
		"list":list,
		"total":total,
	}
	respBytes,_ :=json.Marshal(respData)


	// 3、回写缓存，过期时间30-60s，根据业务调整
	// 回写本地缓存+定时过期
	resp := models.Resp{
		Code: 200,
		Msg: "success",
		Data: respData,
	}
	localCache.Store(cackeKey,resp)
	// // 异步定时清理（简单实现）
	go func(k string){
		time.Sleep(cacheTTL)
		localCache.Delete(k)
	}(cackeKey)
	// utils.LocalCache.Set(cackeKey,resp,1)
	// 回写redis缓存
	expire := cacheTTL+ time.Duration(rand.Intn(10) ) * time.Second
	db.Rdb.Set(c,cackeKey,respBytes,expire)
	c.JSON(200,models.Resp{
		Code: 200,
		Msg: "success",
		Data: respData,
	})
}