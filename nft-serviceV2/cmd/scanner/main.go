package main

import (
	"fmt"
	"nft-service/config"
)


func main() {
	// 初始化配置
	config,err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	fmt.Println(config.DBConfig.DBName)


}