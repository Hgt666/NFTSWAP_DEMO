package main

import (
	"fmt"
	"log"
	"net/http/httputil"
	"net/url"
	"nftswap-etcd/pkg/registry"
	"time"

	"os"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App  registry.Service   `yaml:"app"`
	Etcd registry.EtcdConfig `yaml:"etcd"`
}

var globalCfg Config
var targetAddr atomic.Value // 原子存储后端服务地址（简单负载）

func initConfig() {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	_ = yaml.Unmarshal(data, &globalCfg)
}

// 定时从 etcd 拉取服务列表（服务发现）
func watchService(reg *registry.EtcdRegistry) {
	for {
		// 发现 nft-svc 所有实例
		addrs, err := reg.Discover("nft-svc")
		if err != nil || len(addrs) == 0 {
			log.Println("未发现 nft-svc 服务实例")
			continue
		}
		// 简单取第一个实例，后续可扩展轮询/加权负载均衡
		targetAddr.Store(addrs[0])
		// 5 秒刷新一次
		<-time.After(5 * time.Second)
	}
}

// 反向代理中间件
func proxyMiddleware(c *gin.Context) {
	addr := targetAddr.Load().(string)
	target, err := url.Parse("http://" + addr)
	if err != nil {
		c.JSON(500, gin.H{"msg": "服务不可用"})
		return
	}
	// 转发请求到后端服务
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	initConfig()

	// 1. 连接 etcd
	reg, err := registry.NewEtcdRegistry(globalCfg.Etcd)
	if err != nil {
		log.Fatal("连接 etcd 失败:", err)
	}
	defer reg.Close()

	// 2. 启动协程，持续发现服务
	go watchService(reg)

	// 3. 启动网关
	r := gin.New()
	// 所有请求统一转发到 nft-svc
	r.Any("/api/*path", proxyMiddleware)

	log.Printf("网关启动: %s:%s", globalCfg.App.Addr, globalCfg.App.Port)
	addr :=fmt.Sprintf("%s:%d",globalCfg.App.Addr,globalCfg.App.Port)
	_ = r.Run(addr)

}