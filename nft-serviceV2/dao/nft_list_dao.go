package dao

import (
	"nft-service/db"
	"nft-service/models"
)

// GetNftList 分页查询 NFT 挂单列表
func GetNftList(page, pageSize int) ([]models.NftListDemo, int64, error) {
	var list []models.NftListDemo
	var total int64
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}

	// 总条数
	err := db.DB.Model(&models.NftList{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (page - 1) * pageSize
	// err = db.DB.Order("created_at desc").Limit(pageSize).Offset(offset).Find(&list).Error
	// if err != nil {
	// 	return nil, 0, err
	// }
	// 只查询需要的字段
	err = db.DB.Model(&models.NftList{}).
	Select( "nft_token_id", "price", "created_at").Order("created_at desc").
	Limit(pageSize).Offset(offset).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
