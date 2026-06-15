package models


type NftBid struct {
	ID         uint64  `gorm:"primary_key"`
	ChainId    string  `gorm:"column:chain_id;size:128"`
	NftContract string  `gorm:"column:nft_contract;size:128;uniqueIndex:idx_contract_tokenId"`
	NftTokenId string  `gorm:"column:nft_token_id;size:10;uniqueIndex:idx_contract_tokenId"`
	Price      string `gorm:"column:price;comment:单位ETH"`
	Buyer      string  `gorm:"column:buyer;size:128"`
}
