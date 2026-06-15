package util

import "fmt"

// GenNftListKey 生成NFT列表缓存key
func GenNftListKey(page, pageSize int) string {
	return fmt.Sprintf("nft:list:p%d_s%d", page, pageSize)
}