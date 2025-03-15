package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"crypto-gateway/utils"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"

	"github.com/joho/godotenv"
)

func main() {
	// 加载 .env 文件
	_ = godotenv.Load()

	// 从环境变量中获取 URL
	nodeURL := os.Getenv("ETH_NODE_URL")
    if nodeURL == "" {
        log.Fatal("Error: ETH_NODE_URL is not set in .env or environment variables")
    }

	// 连接主节点
	client, err := ethclient.Dial(nodeURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Ethereum via Infura (loaded from .env)")

	// 1. 定义一个账户地址
    account := common.HexToAddress("0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045")

    // 2. 查询余额 (nil 代表最新区块)
    balance, err := client.BalanceAt(context.Background(), account, nil)
    if err != nil {
        log.Fatal(err)
    }

    // 3. 使用我们的工具进行转换
    ethValue := utils.WeiToEther(balance)

    fmt.Printf("Address: %s\n", account.Hex())
    fmt.Printf("Balance (Wei): %s\n", balance.String())
    // %.4f 保留4位小数
    fmt.Printf("Balance (ETH): %.4f ETH\n", ethValue)
}