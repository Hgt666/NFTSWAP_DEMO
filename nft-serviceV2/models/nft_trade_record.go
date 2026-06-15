package models

type NftTradeRecord struct {
	ID  uint64 `gorm:"primary_key"`
	ChainID string `gorm:"column:chain_id;size:10;uniqueIndex:idx_chainId_nftTokenId_nftContract"` // 链ID，如1（以太坊）等
	NftTokenID string `gorm:"column:nft_token_id;size:10;uniqueIndex:idx_chainId_nftTokenId_nftContract"`
	NftContract string `gorm:"column:nft_contract;size:255;uniqueIndex:idx_chainId_nftTokenId_nftContract"`
	Buyer string `gorm:"column:buyer;size:255"`
	Seller string `gorm:"column:seller;size:255"`
	Price string `gorm:"column:price;size:255"`
	TradeTime string `gorm:"column:trade_time"`
	Status string `gorm:"column:status"` // 交易状态，如"pending"、"completed"等
	TradeType uint8 `gorm:"column:trade_type;comment:1-限价成交 2-市价成交 3-卖家接受出价"` // 交易类型，如"buy"、"sell"等
	TradeID string `gorm:"column:trade_id"` // 交易ID，如订单号等
	TradeHash string `gorm:"column:trade_hash"` // 交易哈希，如以太坊交易哈希等
}