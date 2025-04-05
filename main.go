package main

import (
    "fmt"
    "log"
    "math/big"
    "os"
    "strings"
    
    "crypto-gateway/tx"

    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    
    network := strings.ToLower(os.Getenv("CURRENT_NETWORK"))

    var nodeURL, privKey string

    //根据模式选择配置
    switch network {
    case "sepolia":
        nodeURL = os.Getenv("SEPOLIA_URL")
        privKey = os.Getenv("SEPOLIA_PRIVATE_KEY")
        fmt.Println("🟢 Environment: Sepolia Testnet")
    case "mainnet":
        nodeURL = os.Getenv("MAINNET_URL")
        privKey = os.Getenv("MAINNET_PRIVATE_KEY")
        fmt.Println("⚠️  WARNING: Environment: Ethereum MAINNET")
    default:
        log.Fatalf("Unknown network: %s. Please check .env file.", network)
    }

    // 检查配置是否读取成功
    if nodeURL == "" || privKey == "" {
        log.Fatalf("CRITICAL: Missing configuration for %s network", network)
    }

    // 连接节点
    client, err := ethclient.Dial(nodeURL)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Connected to Node: %s\n", nodeURL)

    // --- 真实转账逻辑 ---
    // 只有在 Sepolia 模式下才自动执行转账，防止主网误操作
    if network == "sepolia" {
        fmt.Println("🚀 Initiating Sepolia transaction...")

        // 转账给 V神
        toAddress := "0xd8dA6BF26964aF9D7eEd9e03E53415D37aA96045"
        
        // 金额: 0.0001 ETH (避免浪费测试币)
        amount := big.NewInt(100000000000000) 

        // 发送！
        txHash := tx.SendETH(client, privKey, toAddress, amount)

        if txHash != "" {
            fmt.Printf("🎉 Transaction Success!\nView on Etherscan: https://sepolia.etherscan.io/tx/%s\n", txHash)
        }
    } else {
        fmt.Println("🛑 Mainnet mode detected. Transaction skipped for safety.")
    }
}