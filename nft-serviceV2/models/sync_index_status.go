package models

type SyncIndexStatus struct {
	ID  uint64 `gorm:"primary_key;auto_increment"`
	NftContractAddress string `gorm:"column:nft_contract_address;size:64"`
	ChainID uint64 `gorm:"column:chain_id"`
	LastIndexedBlock uint64 `gorm:"column:last_indexed_block"`
	LastIndexedTxHash string `gorm:"column:last_indexed_tx_hash;size:128"`
}