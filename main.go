package main

import (
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 连接Cloudflare 的公共节点
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Success! Connected to Ethereum Mainnet")
	_ = client
}