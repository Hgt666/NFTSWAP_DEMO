package dao

import (
	"context"
	"nftswap-etcd/nft-svc/db"
	"nftswap-etcd/nft-svc/models"
)

// GetNftList 分页查询 NFT 挂单列表
// 已配合覆盖索引，只查询必要字段
func GetNftList(ctx context.Context, page, pageSize int) ([]models.NftList, int64, error) {
	var list []models.NftList
	var total int64

	offset := (page - 1) * pageSize

	// 查询总条数
	err := db.DB.WithContext(ctx).
		Model(&models.NftList{}).
		Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 分页查询 + 指定字段 + 排序（命中覆盖索引）
	err = db.DB.WithContext(ctx).
		Select("id, nft_token_id, price, created_at").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&list).Error

	return list, total, err
}

// GetNftByID 根据ID查询单条挂单
func GetNftByID(ctx context.Context, id uint64) (*models.NftList, error) {
	var item models.NftList
	err := db.DB.WithContext(ctx).
		Where("id = ?", id).
		First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

// CreateNft 新增 NFT 挂单
func CreateNft(ctx context.Context, data *models.NftList) error {
	return db.DB.WithContext(ctx).Create(data).Error
}

// UpdateNftStatus 更新挂单状态（成交/下架）
func UpdateNftStatus(ctx context.Context, id uint64, status string) error {
	return db.DB.WithContext(ctx).
		Model(&models.NftList{}).
		Where("id = ?", id).
		Update("status", status).Error
}

// DeleteNft 删除挂单
func DeleteNft(ctx context.Context, id uint64) error {
	return db.DB.WithContext(ctx).
		Where("id = ?", id).
		Delete(&models.NftList{}).Error
}

// ExistByTokenID 校验 TokenID 是否存在（唯一性约束）
func ExistByTokenID(ctx context.Context, tokenID string) (bool, error) {
	var count int64
	err := db.DB.WithContext(ctx).
		Model(&models.NftList{}).
		Where("nft_token_id = ?", tokenID).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}