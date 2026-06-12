package main

import (
	// "nft-service/db"
	"nft-service/db" // 请确保导入正确的包路径
	"go.uber.org/zap" // 请确保导入正确的包路径
	"nft-service/internal"
	"nft-service/sync"
)

func main() {
	// 初始化基础组件
	internal.InitLogger()
	if err := db.InitMysql(); err != nil {
		internal.Logger.Fatal("MySQL init fail", zap.Error(err))
	}
	if err := db.InitRedis(); err != nil {
		internal.Logger.Fatal("Redis init fail", zap.Error(err))
	}

	// 启动区块同步（协程后台运行）
	go sync.StartBlockSync()
	internal.Logger.Info("区块事件同步服务已启动")

	// 启动接口服务
	<-make(chan struct{})
}