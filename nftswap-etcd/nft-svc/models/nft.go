package models

import "time"

// NftList NFT挂单表
type NftList struct {
	ID                uint64    `gorm:"column:id;primaryKey;autoIncrement"`
	NftContractAddress string    `gorm:"column:nft_contract_address"`
	NftTokenID        string    `gorm:"column:nft_token_id"`
	OwnerAddress      string    `gorm:"column:owner_address"`
	Price             string    `gorm:"column:price"`
	Status            string    `gorm:"column:status"` // 0未成交 1已成交
	CreatedAt         time.Time `gorm:"column:created_at"`
	UpdatedAt         time.Time `gorm:"column:updated_at"`
	ChainID           string    `gorm:"column:chain_id"`
}

// TableName 指定表名
func (NftList) TableName() string {
	return "nft_lists"
}

// NftListReq 列表查询入参
type NftListReq struct {
	Page     int `form:"page" binding:"required,min=1"`
	PageSize int `form:"page_size" binding:"required,min=1,max=50"`
}

// Resp 统一返回结构
type Resp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}