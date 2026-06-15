package service

import (
	"nft-service/dao"
	"nft-service/models"
)




// 获取 nft列表
func GetNftList(req models.NftListReq) ([]models.NftList,int64, error) {
	// 实现获取 nft列表的逻辑
	return dao.GetNftList(int(req.Page),int(req.PageSize))
}