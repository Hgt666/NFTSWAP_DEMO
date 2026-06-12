package pkg

import (
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
)

const etherDecimals = 18

// WeiToEth wei 转 ETH (保留8位小数)
func WeiToEth(wei *big.Int) float64 {
	div := new(big.Int).Exp(big.NewInt(10), big.NewInt(etherDecimals), nil)
	eth := new(big.Float).Quo(new(big.Float).SetInt(wei), new(big.Float).SetInt(div))
	f, _ := eth.Float64()
	return f
}

// EthToWei ETH 转 wei
func EthToWei(eth float64) *big.Int {
	mul := new(big.Int).Exp(big.NewInt(10), big.NewInt(etherDecimals), nil)
	ethStr := strconv.FormatFloat(eth, 'f', 18, 64)
	wei, _ := new(big.Int).SetString(ethStr, 10)
	return wei.Mul(wei, mul)
}

// HexToAddr 地址格式化
func HexToAddr(hex string) common.Address {
	return common.HexToAddress(hex)
}