package scanner

import (
	"log"
	"nft-service/db"
	"nft-service/models"

	"gorm.io/gorm/clause"
)

// 存储 OrderList数据
func InsertOrderList(orderList *models.NftList) error {
	err :=db.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name:"chain_id"}, {Name:"nft_contract_address"}, {Name:"nft_token_id"}},
		DoNothing: true,
	}).Create(&orderList).Error
	if err != nil {
		log.Println("NftList Insert error:")
		return err
	}
	return nil
}

// 存储 出价数据
func InsertNftBid(orderList *models.NftBid) error {
	err :=db.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name:"nft_contract"}, {Name:"nft_token_id"}},
		DoNothing: true,
	}).Create(&orderList).Error
	if err != nil {
		log.Println("NftBid Insert error:")
		return err
	}
	return nil
}

// 存储成交数据
func InsertMatchRecord(orderList *models.NftTradeRecord) error {
	err :=db.DB.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name:"chain_id"}, {Name:"nft_token_id"},{Name:"nft_contract"}},
		DoNothing: true,
	}).Create(&orderList).Error
	if err != nil {
		log.Println("MatchRecord Insert error:")
		return err
	}
	return nil
}