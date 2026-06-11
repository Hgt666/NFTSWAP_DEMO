package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "./nft_data.db"
var db *sql.DB

func initDB() {
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	db = conn
}

func main() {
	initDB()
	r := gin.Default()

	// 接口分组
	api := r.Group("/api/nft")
	{
		// 1. 查询单个NFT挂单信息
		api.GET("/list/:tokenId", getListOrder)
		// 2. 查询NFT所有出价
		api.GET("/bid/:tokenId", getBidRecords)
		// 3. 查询NFT成交记录
		api.GET("/trade/:tokenId", getTradeRecords)
		// 4. 分页查询有效挂单
		api.GET("/list/all", getActiveList)
	}

	log.Println("接口服务启动 :8080")
	r.Run(":8080")
}

// 通用返回结构体
type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 1. 查询单个NFT挂单
func getListOrder(c *gin.Context) {
	tokenId := c.Param("tokenId")
	var (
		nftContract, seller, price string
		status                    int
	)
	err := db.QueryRow(`
	SELECT nft_contract,seller,price,status FROM nft_list_order WHERE token_id=?`, tokenId).
		Scan(&nftContract, &seller, &price, &status)
	if err != nil {
		c.JSON(http.StatusOK, Resp{Code: 404, Msg: "暂无挂单", Data: nil})
		return
	}

	c.JSON(http.StatusOK, Resp{
		Code: 200,
		Msg:  "success",
		Data: gin.H{
			"nft_contract": nftContract,
			"token_id":     tokenId,
			"seller":       seller,
			"price_wei":    price,
			"status":       status,
		},
	})
}

// 2. 查询出价记录
func getBidRecords(c *gin.Context) {
	tokenId := c.Param("tokenId")
	rows, err := db.Query(`
	SELECT bidder,bid_price,status FROM nft_bid_record WHERE token_id=? ORDER BY created_at DESC`, tokenId)
	if err != nil {
		c.JSON(http.StatusOK, Resp{Code: 500, Msg: "查询失败", Data: nil})
		return
	}
	defer rows.Close()

	var list []gin.H
	for rows.Next() {
		var bidder, price string
		var status int
		_ = rows.Scan(&bidder, &price, &status)
		list = append(list, gin.H{
			"bidder":    bidder,
			"bid_price": price,
			"status":    status,
		})
	}

	c.JSON(http.StatusOK, Resp{Code: 200, Msg: "success", Data: list})
}

// 3. 查询成交记录
func getTradeRecords(c *gin.Context) {
	tokenId := c.Param("tokenId")
	rows, err := db.Query(`
	SELECT seller,buyer,trade_price,trade_type FROM nft_trade_record WHERE token_id=?`, tokenId)
	if err != nil {
		c.JSON(http.StatusOK, Resp{Code: 500, Msg: "查询失败", Data: nil})
		return
	}
	defer rows.Close()

	var list []gin.H
	for rows.Next() {
		var seller, buyer, price string
		var tType int
		_ = rows.Scan(&seller, &buyer, &price, &tType)
		list = append(list, gin.H{
			"seller":      seller,
			"buyer":       buyer,
			"trade_price": price,
			"trade_type":  tType,
		})
	}
	c.JSON(http.StatusOK, Resp{Code: 200, Msg: "success", Data: list})
}

// 4. 分页查询所有有效挂单
func getActiveList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))
	offset := (page - 1) * size

	rows, err := db.Query(`
	SELECT token_id,seller,price FROM nft_list_order WHERE status=1 LIMIT ? OFFSET ?`, size, offset)
	if err != nil {
		c.JSON(http.StatusOK, Resp{Code: 500, Msg: "查询失败", Data: nil})
		return
	}
	defer rows.Close()

	var list []gin.H
	for rows.Next() {
		var tid, seller, price string
		_ = rows.Scan(&tid, &seller, &price)
		list = append(list, gin.H{
			"token_id": tid,
			"seller":   seller,
			"price":    price,
		})
	}

	c.JSON(http.StatusOK, Resp{Code: 200, Msg: "success", Data: list})
}