package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"nft-service/config"
	"nft-service/db"
	"nft-service/internal"
	"nft-service/pkg"
	"go.uber.org/zap"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func main() {
	internal.InitLogger()
	if err := db.InitMysql(); err != nil {
		internal.Logger.Fatal("MySQL初始化失败", zap.Error(err))
	}
	if err := db.InitRedis(); err != nil {
		internal.Logger.Fatal("Redis初始化失败", zap.Error(err))
	}

	r := gin.Default()
	// 全局中间件
	r.Use(internal.AuthMiddleware())
	r.Use(internal.RateLimitMiddleware())

	api := r.Group("/api/nft")
	{
		api.GET("/list/:tokenId", getListOrder)
		api.GET("/bid/:tokenId", getBidRecords)
		api.GET("/trade/:tokenId", getTradeRecords)
		api.GET("/list/all", getActiveList)
		api.GET("/meta/:tokenId", getNFTMeta) // IPFS元数据
	}

	internal.Logger.Info("接口服务启动", zap.String("addr", config.APIListen))
	r.Run(config.APIListen)
}

func getListOrder(c *gin.Context) {
	tid := c.Param("tokenId")
	var nc, seller, pw string
	var pe float64
	var status int

	err := db.MysqlDB.QueryRow(`
	SELECT nft_contract,seller,price_wei,price_eth,status FROM nft_list_order WHERE token_id=?`, tid).
		Scan(&nc, &seller, &pw, &pe, &status)
	if err != nil {
		c.JSON(http.StatusOK, Resp{Code: 404, Msg: "无挂单数据"})
		return
	}
	c.JSON(http.StatusOK, Resp{Code: 200, Msg: "success", Data: gin.H{
		"nft_contract": nc, "token_id": tid, "seller": seller,
		"price_wei": pw, "price_eth": pe, "status": status,
	}})
}

func getBidRecords(c *gin.Context) {
	tid := c.Param("tokenId")
	rows, err := db.MysqlDB.Query(`
	SELECT bidder,bid_price_wei,bid_price_eth,status FROM nft_bid_record WHERE token_id=?`, tid)
	if err != nil {
		c.JSON(http.StatusOK, Resp{Code: 500, Msg: "查询失败"})
		return
	}
	defer rows.Close()

	var list []gin.H
	for rows.Next() {
		var bidder, pw string
		var pe float64
		var status int
		_ = rows.Scan(&bidder, &pw, &pe, &status)
		list = append(list, gin.H{
			"bidder": bidder, "price_wei": pw, "price_eth": pe, "status": status,
		})
	}
	c.JSON(http.StatusOK, Resp{Code: 200, Msg: "success", Data: list})
}

func getTradeRecords(c *gin.Context) {
	tid := c.Param("tokenId")
	rows, err := db.MysqlDB.Query(`
	SELECT seller,buyer,trade_price_wei,trade_price_eth,trade_type FROM nft_trade_record WHERE token_id=?`, tid)
	if err != nil {
		c.JSON(http.StatusOK, Resp{Code: 500, Msg: "查询失败"})
		return
	}
	defer rows.Close()

	var list []gin.H
	for rows.Next() {
		var seller, buyer, pw string
		var pe float64
		var t int
		_ = rows.Scan(&seller, &buyer, &pw, &pe, &t)
		list = append(list, gin.H{
			"seller": seller, "buyer": buyer, "price_wei": pw, "price_eth": pe, "type": t,
		})
	}
	c.JSON(http.StatusOK, Resp{Code: 200, Msg: "success", Data: list})
}

func getActiveList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	offset := (page - 1) * size

	rows, err := db.MysqlDB.Query(`
	SELECT token_id,seller,price_wei,price_eth FROM nft_list_order WHERE status=1 LIMIT ? OFFSET ?`, size, offset)
	if err != nil {
		c.JSON(http.StatusOK, Resp{Code: 500, Msg: "查询失败"})
		return
	}
	defer rows.Close()

	var list []gin.H
	for rows.Next() {
		var tid, seller, pw string
		var pe float64
		_ = rows.Scan(&tid, &seller, &pw, &pe)
		list = append(list, gin.H{
			"token_id": tid, "seller": seller, "price_wei": pw, "price_eth": pe,
		})
	}
	c.JSON(http.StatusOK, Resp{Code: 200, Msg: "success", Data: list})
}

// 获取NFT IPFS元数据
func getNFTMeta(c *gin.Context) {
	tid := c.Param("tokenId")
	var nc, cid string
	err := db.MysqlDB.QueryRow(`
	SELECT nft_contract,ipfs_cid FROM nft_metadata WHERE token_id=?`, tid).Scan(&nc, &cid)
	if err != nil {
		c.JSON(http.StatusOK, Resp{Code: 404, Msg: "无元数据"})
		return
	}

	meta, err := pkg.FetchIPFSMeta(cid)
	if err != nil {
		c.JSON(http.StatusOK, Resp{Code: 500, Msg: "IPFS拉取失败"})
		return
	}

	c.JSON(http.StatusOK, Resp{Code: 200, Msg: "success", Data: gin.H{
		"nft_contract": nc, "token_id": tid, "ipfs_cid": cid,
		"name": meta.Name, "description": meta.Description, "image": meta.Image,
	}})
}