package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"nft-service/config"
	"nft-service/internal"

	// "nft-service/internal"
	"nft-service/pkg"

	"go.uber.org/zap"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

// "nft-service/internal"


// var MysqlDB *gorm.DB
var MysqlDB *sql.DB


func InitMysql() error {
	db, err := sql.Open("mysql", config.MySQLDSN)
	// db,err := gorm.Open(mysql.Open(config.MySQLDSN),&gorm.Config{})
	if err != nil {
		return err
	}
	// db.SetMaxOpenConns(20)
	// db.SetMaxIdleConns(10)
	MysqlDB = db
	return nil
}

// GetLastSyncBlock 获取最后已处理区块（断点续传）
func GetLastSyncBlock(syncKey string) (uint64, error) {
	var num uint64
	err := MysqlDB.QueryRow("SELECT last_block_num FROM sync_block_point WHERE sync_key = ?", syncKey).Scan(&num)
	return num, err
}

// UpdateSyncBlock 更新同步区块断点
func UpdateSyncBlock(syncKey string, blockNum uint64) error {
	_, err := MysqlDB.Exec("UPDATE sync_block_point SET last_block_num = ? WHERE sync_key = ?", blockNum, syncKey)
	return err
}

// InsertListOrder 插入挂单
func InsertListOrder(nftAddr, seller string, tokenID uint64, priceWei string, priceEth float64, blockNum uint64) {
	_, err := MysqlDB.Exec(`
	INSERT IGNORE INTO nft_list_order(nft_contract,token_id,seller,price_wei,price_eth,block_number)
	VALUES(?,?,?,?,?,?)`, nftAddr, tokenID, seller, priceWei, priceEth, blockNum)
	if err != nil {
		// internal.Logger.Error("插入挂单失败", zap.Error(err))
		// internal.Logger.Error("插入挂单失败", zap.Error(err))
		internal.Logger.Error("插入挂单失败", zap.Error(err))
	}
}

// UpdateListPrice 更新挂单价格
func UpdateListPrice(nftAddr string, tokenID uint64, priceWei string, priceEth float64) {
	_, err := MysqlDB.Exec(`
	UPDATE nft_list_order SET price_wei=?,price_eth=? WHERE nft_contract=? AND token_id=?`,
		priceWei, priceEth, nftAddr, tokenID)
	if err != nil {
		internal.Logger.Error("更新挂单价格失败", zap.Error(err))
	}
}

// CancelListOrder 取消挂单
func CancelListOrder(nftAddr string, tokenID uint64) {
	_, err := MysqlDB.Exec(`
	UPDATE nft_list_order SET status=0 WHERE nft_contract=? AND token_id=?`, nftAddr, tokenID)
	if err != nil {
		internal.Logger.Error("取消挂单失败", zap.Error(err))
	}
}

// InsertBid 插入出价
func InsertBid(nftAddr, bidder string, tokenID uint64, priceWei string, priceEth float64, blockNum uint64) {
	_, err := MysqlDB.Exec(`
	INSERT INTO nft_bid_record(nft_contract,token_id,bidder,bid_price_wei,bid_price_eth,block_number)
	VALUES(?,?,?,?,?,?)`, nftAddr, tokenID, bidder, priceWei, priceEth, blockNum)
	if err != nil {
		internal.Logger.Error("插入出价失败", zap.Error(err))
	}
}

// CancelBid 取消出价
func CancelBid(nftAddr string, tokenID uint64) {
	_, err := MysqlDB.Exec(`
	UPDATE nft_bid_record SET status=0 WHERE nft_contract=? AND token_id=?`, nftAddr, tokenID)
	if err != nil {
		internal.Logger.Error("取消出价失败", zap.Error(err))
	}
}

// InsertTrade 插入成交记录
func InsertTrade(nftAddr, seller, buyer, priceWei string, priceEth float64,
	tokenID uint64, tradeType uint8, blockNum uint64) {
	_, err := MysqlDB.Exec(`
	INSERT INTO nft_trade_record(nft_contract,token_id,seller,buyer,trade_price_wei,trade_price_eth,trade_type,block_number)
	VALUES(?,?,?,?,?,?,?,?)`, nftAddr, tokenID, seller, buyer, priceWei, priceEth, tradeType, blockNum)
	if err != nil {
		internal.Logger.Error("插入成交记录失败", zap.Error(err))
	}
	// 成交关闭挂单
	CancelListOrder(nftAddr, tokenID)
}

// UpsertMetadata 写入NFT元数据
func UpsertMetadata(nftAddr string, tokenID uint64, cid string, meta *pkg.NFTMetadata) {
	_, err := MysqlDB.Exec(`
	INSERT INTO nft_metadata(nft_contract,token_id,ipfs_cid,name,description,image)
	VALUES(?,?,?,?,?,?)
	ON DUPLICATE KEY UPDATE ipfs_cid=?,name=?,description=?,image=?`,
		nftAddr, tokenID, cid, meta.Name, meta.Description, meta.Image,
		cid, meta.Name, meta.Description, meta.Image)
	if err != nil {
		internal.Logger.Error("写入元数据失败", zap.Error(err))
	}
}