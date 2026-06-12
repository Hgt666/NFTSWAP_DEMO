package config

import "time"

// 以太坊配置
var (
	EthRPC       = "http://127.0.0.1:8545"
	MarketAddr   = "0xe7f1725E7734CE288F8367e1Bb143E90bb3F0512"
	NFTContract  = "0x5FbDB2315678afecb367f032d93F642f64180aa3"
	SyncKey      = "nft_market_event"
	SyncInterval = 1 * time.Second // 区块轮询间隔
)

// MySQL
var (
	MySQLDSN = "root:rootpasswd@tcp(127.0.0.1:3307)/test?charset=utf8mb4&parseTime=True&loc=Local"
)

// Redis
var (
	RedisAddr = "127.0.0.1:6379"
)

// 接口配置
var (
	APIListen = ":8080"
	APIToken  = "nft-service-2026" // 接口鉴权Token
)

// IPFS
var (
	IPFSGateway = "https://ipfs.io/ipfs/"
)