package main

import (
    "fmt"
    "log"
    "math/big"
    "os"
    
    "crypto-gateway/tx"

    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/joho/godotenv"
)

func main() {
    _ = godotenv.Load()
    nodeURL := os.Getenv("ETH_NODE_URL")
    
    client, err := ethclient.Dial(nodeURL)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to Ethereum")

    // --- 模拟转账测试 ---
    
    // 1. 一个随机生成的测试私钥 (空钱包，用于测试代码逻辑)
    testPrivKey := "fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19"
    
    // 2. 转给一个地址 (V神)
    toAddress := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
    
    // 3. 转账金额: 0.01 ETH (10000000000000000 Wei)
    amount := big.NewInt(10000000000000000)

    fmt.Println("Attempting to send transaction...")
    
    // 调用我们的转账模块
    tx.SendETH(client, testPrivKey, toAddress, amount)
}