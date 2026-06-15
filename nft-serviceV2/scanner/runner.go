package scanner

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"nft-service/config"
	"nft-service/db"
	"nft-service/models"
	"time"

	"github.com/ethereum/go-ethereum/common"

	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum"
	"gorm.io/gorm"
)

//事件签名

var (
	OrderListedSig = common.HexToHash("0xb4aa437e32593e335dd4eb6069bc5e3225445ac88d9e797c5b9b99c404d9dcdd")
	BidPlacedSig   = common.HexToHash("0xdd49bbb40d47a514dddcd458e9718364143bc24a0cca58439ee6f4f45e4ce10d")
	MatchSig       = common.HexToHash("0x7e769bd821f1c5a635e0bd4c3654e8630e529600e771ec4b9018b0c0bc5a928d")
)

// Run 扫描服务入口，接收全局根 Ctx
func Run(ctx context.Context) {
	// 初始化链客户端
	if err := InitEthClient(); err != nil {
		log.Println("err", err)
		log.Fatal("初始化 eth客户端失败:%w")
	}
	defer CloseEthClient()

	cfg := config.GlobalConfig.Chain
	scanInterval := time.Duration(cfg.ScanInterval) * time.Second
	// 定时扫描
	ticker := time.NewTicker(scanInterval)
	defer ticker.Stop()

	log.Printf("开始扫描服务，扫描间隔:%dms, 目标合约:%s", scanInterval, cfg.TargetContract)

	// 主循环扫描
	for {
		select {
		case <-ctx.Done():
			log.Println("扫描退出")
			return
		case <-ticker.C: // 每隔 scanInterval 扫描一次
			// 执行扫描逻辑
			if err := scanOnce(ctx, cfg); err != nil {
				log.Println("err", err)
				return
			}
			// default:

			// 	time.Sleep(scanInterval)
		}
	}
}

// 单次扫描逻辑
func scanOnce(ctx context.Context, cfg config.ChainConfig) error {
	// 获取链上最新高度
	latestBlock, err := GetLatesBlockHeight(ctx)
	if err != nil {
		log.Println("err", err)
		return err
	}
	log.Printf("最新高度:%d", latestBlock)
	var lastSyncIndexStatus models.SyncIndexStatus
	err = db.DB.Model(&models.SyncIndexStatus{}).
		Where("nft_contract_address = ? and chain_id = ?", cfg.TargetContract, cfg.ChainId).First(&lastSyncIndexStatus).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("err", err)
		return err
	}
	// 上次索引高度
	lastSyncBlock := lastSyncIndexStatus.LastIndexedBlock
	log.Printf("上次索引高度:%d", lastSyncBlock)
	// 没有新区块，直接返回
	if lastSyncBlock >= latestBlock {
		return nil
	}

	// 计算本次扫描范围
	endBlock := lastSyncBlock + cfg.BatchSize
	if endBlock > latestBlock {
		endBlock = latestBlock
	}
	log.Printf("本次扫描范围:%d-%d", lastSyncBlock+1, endBlock)

	// 获取 endBlock的BlockHash
	blockHash, err := ethClient.BlockByNumber(ctx, big.NewInt(int64(endBlock)))
	if err != nil {
		log.Println("err", err)
		return err
	}

	// 扫描本次区块范围，日志解析，去重，入库
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(lastSyncBlock + 1)),
		ToBlock:   big.NewInt(int64(endBlock)),
		Addresses: []common.Address{common.HexToAddress(cfg.TargetContract)},
	}
	logs, err := ethClient.FilterLogs(ctx, query)
	if err != nil {
		log.Println("err", err)
		return err
	}

	// 按topics[0]分发日志解析处理
	for _, log := range logs {
		// 解析日志
		switch log.Topics[0] {
		case OrderListedSig:
			ParseOrderListEvent(log)
		case BidPlacedSig:
			ParseBidPlacedEvent(log)
		case MatchSig:
			ParseMatchEvent(log) // 这里需要实现 MatchEvent 的解析逻辑
		default:
			fmt.Println("未知日志")
		}
	}

	// 更新索引状态
	// if err:= db.DB.Model(&models.SyncIndexStatus{}).
	// Where("nft_contract_address = ? and chain_id = ?", cfg.TargetContract, cfg.ChainId).
	// Update("last_indexed_block", endBlock).Error; err != nil {
	// 	log.Println("err",err)
	// 	return err
	// }
	if err := db.DB.Model(&models.SyncIndexStatus{}).
		Where("nft_contract_address = ? and chain_id = ?", cfg.TargetContract, cfg.ChainId).Updates(models.SyncIndexStatus{
		LastIndexedBlock: endBlock,
		// LastIndexedTxHash: string(blockHash.Hash().Bytes()),
		LastIndexedTxHash: blockHash.Hash().Hex(), // 这里需要将 blockHash 转换为字符串
	}).Error; err != nil {
		log.Println("err", err)
		return err
	}
	log.Printf("更新索引状态成功，索引高度:%d, txHash:%s", endBlock, blockHash.Hash().Hex())

	return nil

}
