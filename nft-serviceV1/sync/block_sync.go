package sync

import (
	"context"
	"math/big"
	"time"

	"nft-service/config"
	"nft-service/contract"
	"nft-service/db"
	"nft-service/internal"
	"nft-service/pkg"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"
)

func StartBlockSync() {
	client, err := ethclient.Dial(config.EthRPC)
	if err != nil {
		internal.Logger.Fatal("连接ETH RPC失败", zap.Error(err))
	}
	defer client.Close()

	marketAddr := common.HexToAddress(config.MarketAddr)
	marketContract, err := contract.NewNftMarket(marketAddr, client)
	if err != nil {
		internal.Logger.Fatal("绑定合约失败", zap.Error(err))
	}

	for {
		// 1. 获取当前链最高区块
		latestBlock, err := client.BlockNumber(context.Background())
		if err != nil {
			internal.SimpleAlert("获取最新区块失败", err.Error())
			time.Sleep(config.SyncInterval)
			continue
		}

		// 2. 获取上次处理到的区块（断点）
		lastSyncBlock, err := db.GetLastSyncBlock(config.SyncKey)
		if err != nil {
			internal.SimpleAlert("读取同步断点失败", err.Error())
			time.Sleep(config.SyncInterval)
			continue
		}

		// 已经同步到最新，等待
		if lastSyncBlock >= latestBlock {
			time.Sleep(config.SyncInterval)
			continue
		}

		// 3. 逐区块拉取日志
		start := lastSyncBlock + 1
		for blockNum := start; blockNum <= latestBlock; blockNum++ {
			filter := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(blockNum)),
				ToBlock:   big.NewInt(int64(blockNum)),
				Addresses: []common.Address{marketAddr},
			}

			logs, err := client.FilterLogs(context.Background(), filter)
			if err != nil {
				internal.SimpleAlert("拉取区块日志失败", err.Error())
				continue
			}

			// 解析每条日志
			for _, vLog := range logs {
				handleLog(marketContract, vLog, blockNum)
			}

			// 4. 更新断点
			_ = db.UpdateSyncBlock(config.SyncKey, blockNum)
		}

		time.Sleep(config.SyncInterval)
	}
}

func handleLog(market *contract.NftMarket, vLog types.Log, blockNum uint64) {
	nftAddr := config.NFTContract

	// 挂单事件
	if e, err := market.ParseOrderListed(vLog); err == nil {
		weiStr := e.Price.String()
		ethVal := pkg.WeiToEth(e.Price)
		db.InsertListOrder(nftAddr, e.Seller.Hex(), e.TokenId.Uint64(), weiStr, ethVal, blockNum)
		internal.Logger.Info("解析挂单事件", zap.Uint64("block", blockNum), zap.Uint64("tokenId", e.TokenId.Uint64()))
		return
	}

	// 修改挂单
	if e, err := market.ParseOrderEdited(vLog); err == nil {
		weiStr := e.NewPrice.String()
		ethVal := pkg.WeiToEth(e.NewPrice)
		db.UpdateListPrice(nftAddr, e.TokenId.Uint64(), weiStr, ethVal)
		return
	}

	// 取消挂单
	if e, err := market.ParseOrderCancelled(vLog); err == nil {
		db.CancelListOrder(nftAddr, e.TokenId.Uint64())
		return
	}

	// 出价
	if e, err := market.ParseBidPlaced(vLog); err == nil {
		weiStr := e.Price.String()
		ethVal := pkg.WeiToEth(e.Price)
		db.InsertBid(nftAddr, e.Bidder.Hex(), e.TokenId.Uint64(), weiStr, ethVal, blockNum)
		return
	}

	// 取消出价
	if e, err := market.ParseBidCancelled(vLog); err == nil {
		db.CancelBid(nftAddr, e.TokenId.Uint64())
		return
	}

	// 成交
	if e, err := market.ParseMatchSuccess(vLog); err == nil {
		weiStr := e.Price.String()
		ethVal := pkg.WeiToEth(e.Price)
		db.InsertTrade(nftAddr, e.Seller.Hex(), e.Buyer.Hex(), weiStr, ethVal,
			e.TokenId.Uint64(), uint8(e.TradeType), blockNum)
		return
	}
}