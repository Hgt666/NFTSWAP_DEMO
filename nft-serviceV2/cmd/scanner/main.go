package main

import (
	"context"
	"log"
	"nft-service/config"
	"nft-service/db"
	"nft-service/mq"
	"nft-service/scanner"
)

var (
	Ctx = context.Background()
)

func main() {
	// 初始化配置
	err := config.InitConfig()
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}
	log.Println("配置初始化成功")
	// 初始化mysql
	err = db.InitMySQL()
	if err != nil {
		log.Fatalf("Failed to initialize mysql: %v", err)
	}
	log.Println("mysql初始化成功")

	// 初始化redis
	err = db.InitRedis(Ctx)
	if err != nil {
		log.Fatalf("Failed to initialize redis: %v", err)
	}
	log.Println("redis初始化成功")

	// 初始化 mq
	err =mq.InitRabbitMQ()
	if err != nil {
		log.Fatalf("Failed to initialize mq: %v", err)
	}
	log.Println("mq初始化成功")

	// 启动扫链服务
	scanner.Run(Ctx)


}
