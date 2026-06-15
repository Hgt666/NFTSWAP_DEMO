package utils

import (
	"fmt"
)

// wei转成 ETH
func WeiToEth(wei uint64) string {
	return fmt.Sprintf("%.18f", float64(wei)/1e18)
}