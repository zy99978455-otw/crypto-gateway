package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ethereum/go-ethereum/ethclient"
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

	// 获取最新的区块号
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Latest Block Number: %d\n", blockNumber)
}