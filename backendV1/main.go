package main

import (
	"context"
	"log"

	// "math/big"

	"github.com/ethereum/go-ethereum"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"demo/contract"
	"demo/db"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// 配置项
const (
	// RPC 地址：本地Anvil / Sepolia测试网
	rpcURL      = "ws://127.0.0.1:8545"
	marketAddr  = "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512" // 交易市场合约地址
	nftContract = "0x5FbDB2315678afecb367f032d93F642f64180aa3"  // NFT合约地址
	sqlitePath  = "./nft_data.db"
)

func main() {
	// 1. 初始化数据库
	if err := db.Init(sqlitePath); err != nil {
		log.Fatalf("数据库初始化失败: %v", err)
	}
	log.Println("✅ 数据库初始化完成")

	// 2. 连接ETH节点
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("连接RPC节点失败: %v", err)
	}
	defer client.Close()
	log.Println("✅ 连接ETH RPC节点成功")

	// 3. 绑定合约实例
	// marketContract, err := contract.NewNFTTradeMarket(common.HexToAddress(marketAddr), client)
	marketContract, err := contract.NewNftMarket(common.HexToAddress(marketAddr), client)
	if err != nil {
		log.Fatalf("绑定合约实例失败: %v", err)
	}
	log.Println("✅ 绑定合约实例成功")
	

	// // 4. 构造日志过滤：监听当前合约所有事件
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(marketAddr)},
	}
	
	

	logsChan := make(chan types.Log)
	// 订阅区块日志
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logsChan)
	if err != nil {
		log.Fatalf("订阅事件失败: %v", err)
	}
	log.Println("✅ 开始监听合约事件...")

	// 5. 循环消费事件
	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("事件订阅中断: %v", err)
		case vLog := <-logsChan:
			handleEvent(marketContract, vLog)
		}
	}
}

// 统一解析所有事件
func handleEvent(contract *contract.NftMarket, vLog types.Log) {
	blockNum := vLog.BlockNumber
	nftAddr := nftContract

	// 1. 解析 OrderListed 挂单事件
	eventList, err := contract.ParseOrderListed(vLog)
	if err == nil {
		log.Printf("[挂单] tokenId:%d, 卖家:%s, 价格:%s",
			eventList.TokenId.Uint64(), eventList.Seller.Hex(), eventList.Price.String())
		db.InsertListOrder(nftAddr, eventList.Seller.Hex(), eventList.TokenId.Uint64(),
			eventList.Price.String(), blockNum)
		return
	}

	// 2. 解析 OrderEdited 修改挂单
	eventEdit, err := contract.ParseOrderEdited(vLog)
	if err == nil {
		log.Printf("[改价] tokenId:%d, 新价格:%s", eventEdit.TokenId.Uint64(), eventEdit.NewPrice.String())
		db.UpdateListOrderPrice(nftAddr, eventEdit.TokenId.Uint64(), eventEdit.NewPrice.String())
		return
	}

	// 3. 解析 OrderCancelled 取消挂单
	eventCancelList, err := contract.ParseOrderCancelled(vLog)
	if err == nil {
		log.Printf("[取消挂单] tokenId:%d", eventCancelList.TokenId.Uint64())
		db.CancelListOrder(nftAddr, eventCancelList.TokenId.Uint64())
		return
	}

	// 4. 解析 BidPlaced 出价
	eventBid, err := contract.ParseBidPlaced(vLog)
	if err == nil {
		log.Printf("[出价] tokenId:%d, 出价人:%s, 金额:%s",
			eventBid.TokenId.Uint64(), eventBid.Bidder.Hex(), eventBid.Price.String())
		db.InsertBidRecord(nftAddr, eventBid.Bidder.Hex(), eventBid.TokenId.Uint64(),
			eventBid.Price.String(), blockNum)
		return
	}

	// 5. 解析 BidCancelled 取消出价
	eventCancelBid, err := contract.ParseBidCancelled(vLog)
	if err == nil {
		log.Printf("[取消出价] tokenId:%d", eventCancelBid.TokenId.Uint64())
		db.CancelBid(nftAddr, eventCancelBid.TokenId.Uint64())
		return
	}

	// 6. 解析 MatchSuccess 成交
	eventTrade, err := contract.ParseMatchSuccess(vLog)
	if err == nil {
		tradeType := uint8(eventTrade.TradeType)
		log.Printf("[成交] tokenId:%d, 卖家:%s, 买家:%s, 类型:%d, 价格:%s",
			eventTrade.TokenId.Uint64(), eventTrade.Seller.Hex(), eventTrade.Buyer.Hex(),
			tradeType, eventTrade.Price.String())
		db.InsertTradeRecord(nftAddr, eventTrade.Seller.Hex(), eventTrade.Buyer.Hex(),
			eventTrade.Price.String(), eventTrade.TokenId.Uint64(), tradeType, blockNum)
		return
	}
}