package db

import (
	"database/sql"
	// "fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Init 初始化SQLite
func Init(dbPath string) error {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(1 * time.Hour)
	DB = db

	// 初始化三张表
	err = createTables()
	return err
}

func createTables() error {
	sqlListOrder := `
	CREATE TABLE IF NOT EXISTS nft_list_order (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nft_contract TEXT NOT NULL,
		token_id INTEGER NOT NULL,
		seller TEXT NOT NULL,
		price TEXT NOT NULL,
		status TINYINT NOT NULL DEFAULT 1,
		block_number INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(nft_contract, token_id)
	);`
	_, err := DB.Exec(sqlListOrder)
	if err != nil {
		return err
	}

	sqlBid := `
	CREATE TABLE IF NOT EXISTS nft_bid_record (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nft_contract TEXT NOT NULL,
		token_id INTEGER NOT NULL,
		bidder TEXT NOT NULL,
		bid_price TEXT NOT NULL,
		status TINYINT NOT NULL DEFAULT 1,
		block_number INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = DB.Exec(sqlBid)
	if err != nil {
		return err
	}

	sqlTrade := `
	CREATE TABLE IF NOT EXISTS nft_trade_record (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nft_contract TEXT NOT NULL,
		token_id INTEGER NOT NULL,
		seller TEXT NOT NULL,
		buyer TEXT NOT NULL,
		trade_price TEXT NOT NULL,
		trade_type TINYINT NOT NULL,
		block_number INTEGER NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = DB.Exec(sqlTrade)
	return err
}

// ========== 挂单相关 ==========
func InsertListOrder(nftAddr, seller string, tokenID uint64, price string, blockNum uint64) {
	_, err := DB.Exec(`
	INSERT OR IGNORE INTO nft_list_order(nft_contract,token_id,seller,price,block_number)
	VALUES(?,?,?,?,?)`, nftAddr, tokenID, seller, price, blockNum)
	if err != nil {
		log.Printf("插入挂单失败: %v", err)
	}
}

func UpdateListOrderPrice(nftAddr string, tokenID uint64, newPrice string) {
	_, err := DB.Exec(`
	UPDATE nft_list_order SET price=? WHERE nft_contract=? AND token_id=?`,
		newPrice, nftAddr, tokenID)
	if err != nil {
		log.Printf("更新挂单价格失败: %v", err)
	}
}

func CancelListOrder(nftAddr string, tokenID uint64) {
	_, err := DB.Exec(`
	UPDATE nft_list_order SET status=0 WHERE nft_contract=? AND token_id=?`,
		nftAddr, tokenID)
	if err != nil {
		log.Printf("取消挂单失败: %v", err)
	}
}

// ========== 出价相关 ==========
func InsertBidRecord(nftAddr, bidder string, tokenID uint64, bidPrice string, blockNum uint64) {
	_, err := DB.Exec(`
	INSERT INTO nft_bid_record(nft_contract,token_id,bidder,bid_price,block_number)
	VALUES(?,?,?,?,?)`, nftAddr, tokenID, bidder, bidPrice, blockNum)
	if err != nil {
		log.Printf("插入出价记录失败: %v", err)
	}
}

func CancelBid(nftAddr string, tokenID uint64) {
	_, err := DB.Exec(`
	UPDATE nft_bid_record SET status=0 WHERE nft_contract=? AND token_id=?`,
		nftAddr, tokenID)
	if err != nil {
		log.Printf("取消出价失败: %v", err)
	}
}

// ========== 成交相关 ==========
func InsertTradeRecord(nftAddr, seller, buyer, price string, tokenID uint64, tradeType uint8, blockNum uint64) {
	_, err := DB.Exec(`
	INSERT INTO nft_trade_record(nft_contract,token_id,seller,buyer,trade_price,trade_type,block_number)
	VALUES(?,?,?,?,?,?,?)`, nftAddr, tokenID, seller, buyer, price, tradeType, blockNum)
	if err != nil {
		log.Printf("插入成交记录失败: %v", err)
	}
	// 成交后关闭挂单
	CancelListOrder(nftAddr, tokenID)
}