package models

import "github.com/ethereum/go-ethereum/common"

// 链上事件
type NFTEvent struct {
	BlockNumber uint64 `json:"blockNumber"`
	BlockHash common.Hash `json:"blockHash"`
	TxHash common.Hash `json:"txHash"`
	From common.Address `json:"from"`
	To common.Address `json:"to"`
	TokenID uint64 `json:"tokenId"`
	EventName string `json:"eventName"`
}