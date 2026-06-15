package service

import (
	"context"
	"nftswap-etcd/nft-svc/dao"
	"nftswap-etcd/nft-svc/models"
)

func GetNftList(ctx context.Context, page, pageSize int) ([]models.NftList, int64, error) {
	return dao.GetNftList(ctx, page, pageSize)
}