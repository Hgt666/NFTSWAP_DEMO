package models

import (

	"time"
)

type NftList struct {
	ID uint `gorm:"primary_key;auto_increment"`
	ChainId string `gorm:"column:chain_id;size:10;index"`
	NftContractAddress string `gorm:"column:nft_contract_address;size:128;uniqueIndex:idx_contract_token"`
	NftTokenId string `gorm:"column:nft_token_id;size:128;uniqueIndex:idx_contract_token"`
	OwnerAddress string `gorm:"column:owner_address;size:128;"`
	Price string `gorm:"column:price;size:128;index"`
	Status string `gorm:"column:status;size:10;comment:0未成交 1已成交"`
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}


type NftListDemo struct {
	NftTokenId string `json:"nft_token_id"`
	Price string `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}


// 分页查询请求参数 
type NftListReq struct {
	Page uint `form:"page" binding:"required,min=1"`
	PageSize uint `form:"page_size" binding:"required,min=1,max=100"`
}


// 统一接口返回格式
type Resp struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}