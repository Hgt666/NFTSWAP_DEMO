package scanner

import (
	"fmt"
	"math/big"
	"nft-service/config"
	"nft-service/models"
	"nft-service/utils"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

// 读取 abi文件
func LoadABI() (abi.ABI, error) {
	// 读取 abi 文件
	abiBytes, err := os.ReadFile("./contract/abi/NFTTradeMarket.abi")
	if err != nil {
		fmt.Println("读取abi文件err", err)
		return abi.ABI{}, err
	}
	// 解析 abi 文件
	abiParser, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		return abi.ABI{}, err
	}
	return abiParser, nil

}

// 解析挂单事件
func ParseOrderListEvent(log types.Log) {
	type OrderListEvent struct {
		Seller      string
		NftContract string
		TokenId     string
		Price       *big.Int
	}
	var orderListEvent OrderListEvent
	// 解析indexed字段
	if len(log.Topics) > 3 {
		seller := log.Topics[1].Hex()
		nftContract := log.Topics[2].Hex()
		tokenId := log.Topics[3].Hex()
		orderListEvent.Seller = seller
		orderListEvent.NftContract = nftContract
		orderListEvent.TokenId = tokenId
	}

	// 通过 ABI 解析事件
	abiParser, err := LoadABI()
	if err != nil {
		fmt.Println(err)
	}
	// 把log解析进结构体
	err = abiParser.UnpackIntoInterface(&orderListEvent, "OrderListed", log.Data)
	if err != nil {
		fmt.Println("解析logerr", err)
	}
	// 入库
	// var nftList models.NftList
	nftList := models.NftList{
		ChainId:            config.GlobalConfig.Chain.ChainId,
		NftContractAddress: orderListEvent.NftContract,
		NftTokenId:         orderListEvent.TokenId,
		OwnerAddress:       orderListEvent.Seller,
		Price:              orderListEvent.Price.String(),
		Status:             "0", // 0未成交 1已成交
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	err = InsertOrderList(&nftList)
	if err != nil {
		fmt.Println("入库err", err)
	}

}

// 解析出价事件
func ParseBidPlacedEvent(log types.Log) {
	type BidPlacedEvent struct {
		Bidder      string
		NftContract string
		TokenId     *big.Int
		Price       *big.Int
	}
	var bidPlacedEvent BidPlacedEvent
	// 解析indexed字段
	if len(log.Topics) > 3 {
		bidder := log.Topics[1].Hex()
		nftContract := log.Topics[2].Hex()
		tokenId := log.Topics[3].Big()
		bidPlacedEvent.Bidder = bidder
		bidPlacedEvent.NftContract = nftContract
		bidPlacedEvent.TokenId = tokenId
	}

	// 通过 ABI 解析事件
	abiParser, err := LoadABI()
	if err != nil {
		fmt.Println(err)
	}
	// 把log解析进结构体
	err = abiParser.UnpackIntoInterface(&bidPlacedEvent, "BidPlaced", log.Data)
	if err != nil {
		fmt.Println("解析logerr", err)
	}
	// 入库
	// var nftList models.NftList
	priceEth := utils.WeiToEth(uint64(bidPlacedEvent.Price.Int64()))
	nftBid := models.NftBid{
		ChainId:     config.GlobalConfig.Chain.ChainId,
		NftContract: bidPlacedEvent.NftContract,
		NftTokenId:  bidPlacedEvent.TokenId.String(),
		Price:       priceEth,
		Buyer:       bidPlacedEvent.Bidder,
	}

	err = InsertNftBid(&nftBid)
	if err != nil {
		fmt.Println("入库err", err)
	}

}

// 解析成交事件
func ParseMatchEvent(log types.Log) {
	type MatchEvent struct {
		Seller      string
		Buyer       string
		NftContract string
		TokenId     *big.Int
		Price       *big.Int
		TradeType   uint8
	}
	var matchEvent MatchEvent
	// 解析indexed字段
	if len(log.Topics) > 3 {
		seller := log.Topics[1].Hex()
		buyer := log.Topics[2].Hex()
		nftContract := log.Topics[3].Hex()
		matchEvent.Seller = seller
		matchEvent.NftContract = nftContract
		matchEvent.Buyer = buyer
	}

	// 通过 ABI 解析事件
	abiParser, err := LoadABI()
	if err != nil {
		fmt.Println(err)
	}
	// 把log解析进结构体
	err = abiParser.UnpackIntoInterface(&matchEvent, "MatchSuccess", log.Data)
	if err != nil {
		fmt.Println("解析logerr", err)
	}
	// 入库
	// var nftList models.NftList
	priceEth := utils.WeiToEth(uint64(matchEvent.Price.Int64()))
	nftTradeRecord := models.NftTradeRecord{
		ChainID:     config.GlobalConfig.Chain.ChainId,
		NftContract: matchEvent.NftContract,
		NftTokenID:  matchEvent.TokenId.String(),
		Buyer:       matchEvent.Buyer,
		Seller:       matchEvent.Seller,
		Price:       priceEth,
		TradeTime:    time.Now().Format("2006-01-02 15:04:05"),
		Status:       "1", // 0未成交 1已成交
		TradeType:    matchEvent.TradeType,
	}

	err = InsertMatchRecord(&nftTradeRecord)
	if err != nil {
		fmt.Println("入库err", err)
	}

}
