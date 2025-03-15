package utils

import (
    "math"
    "math/big"
)

// WeiToEther 将 Wei (整数) 转换为 ETH (浮点数)
// 使用 big.Float 保证高精度计算，防止丢失小数点后的金额
func WeiToEther(wei *big.Int) *big.Float {
    // 创建一个高精度的浮点数变量
    fWei := new(big.Float)

    // 把 Wei (整数) 设置进去
    fWei.SetString(wei.String())

    // 创建除数 (10的18次方)
    ethValue := new(big.Float).Quo(fWei, big.NewFloat(math.Pow10(18)))

    return ethValue
}