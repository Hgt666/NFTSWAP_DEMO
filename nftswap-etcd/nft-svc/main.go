package main

import (
	"log"
	"nftswap-etcd/nft-svc/api"
	"nftswap-etcd/nft-svc/common"
	"nftswap-etcd/nft-svc/db"
	"nftswap-etcd/pkg/cache"
	"nftswap-etcd/pkg/registry"
	"fmt"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"os"
)

// 全局配置结构体
type GlobalConfig struct {
	App        registry.Service     `yaml:"app"`
	Etcd       registry.EtcdConfig  `yaml:"etcd"`
	Mysql      db.MysqlConfig       `yaml:"mysql"`
	Redis      cache.RedisConfig    `yaml:"redis"`
	LocalCache cache.LocalCacheConfig `yaml:"local_cache"`
}

var globalCfg GlobalConfig

// 加载yaml配置
func initConfig() {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("读取配置文件失败:", err)
	}
	if err := yaml.Unmarshal(data, &globalCfg); err != nil {
		log.Fatal("解析配置失败:", err)
	}
}

func main() {
	// 1. 加载配置
	initConfig()

	// 2. 初始化 MySQL / Redis / 本地缓存
	if err := common.InitAll(globalCfg.Mysql, globalCfg.Redis, globalCfg.LocalCache); err != nil {
		log.Fatal("组件初始化失败:", err)
	}

	// 3. 连接etcd并注册服务
	reg, err := registry.NewEtcdRegistry(globalCfg.Etcd)
	if err != nil {
		log.Fatal("连接etcd失败:", err)
	}
	defer reg.Close()
	addr :=fmt.Sprintf("%s:%d",globalCfg.App.Addr,globalCfg.App.Port)
	svc := registry.Service{
		Name: "nft-svc",
		Addr: addr,
	}
	if err := reg.Register(svc, 10); err != nil {
		log.Fatal("服务注册失败:", err)
	}

	// 4. 启动Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	api.RegisterNftRouter(r)

	log.Printf("nft-svc 启动成功，监听 %s", svc.Addr)
	_ = r.Run(svc.Addr)
}