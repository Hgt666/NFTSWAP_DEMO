package scanner

import (
	"context"
	"fmt"
	"nft-service/config"

	"github.com/ethereum/go-ethereum/ethclient"
)

// 负责连接 RPC，拉取区块、日志、交易、封装基础方法

// 全局 ETH 客户端
var ethClient *ethclient.Client


// 初始化 ETH 客户端
func InitEthClient() error {
	// 连接到 ETH RPC 服务器
	cfg := config.GlobalConfig.Chain
	if cfg.RpcUrl == "" {
		return fmt.Errorf("rpc url is empty")
	}
	client,err := ethclient.Dial(cfg.RpcUrl)
	if err != nil {
		return fmt.Errorf("dial rpc failed: %w",err)
	}
	ethClient = client
	return nil
}



func CloseEthClient() {
	if ethClient != nil {
		ethClient.Close()
	}
}


// 获取链上最新区块高度
func GetLatesBlockHeight(ctx context.Context) (uint64,error) {
	header,err := ethClient.HeaderByNumber(ctx,nil)
	if err != nil {
		return 0,err
	}
	return header.Number.Uint64(),nil
}