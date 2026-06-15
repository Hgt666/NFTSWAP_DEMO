package models

import (
	"time"
)

type NftMeta struct {
	ID uint64 `gorm:"primary_key"`
	NftTokenId string `gorm:"column:nft_token_id"`
	Metadata string `gorm:"column:metadata"`
	Owner string `gorm:"column:owner"`
	ContractAddress string `gorm:"column:contract_address"`
	ChainId string `gorm:"column:chain_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}