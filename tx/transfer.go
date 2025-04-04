package tx

import (
    "context"
    "crypto/ecdsa"
    "log"
    "math/big"

    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
)

// SendETH 发送原生 ETH
// client: 节点连接
// privKeyHex: 私钥 (16进制字符串)
// toStr: 接收方地址
// amountWei: 发送金额 (Wei)
func SendETH(client *ethclient.Client, privKeyHex string, toStr string, amountWei *big.Int) string {
    // 1. 私钥处理 (把字符串转为对象)
    privateKey, err := crypto.HexToECDSA(privKeyHex)
    if err != nil {
        log.Fatal("Invalid private key:", err)
    }

    // 2. 推导发送方地址 (Public Key -> Address)
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("Error casting public key to ECDSA")
    }
    fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

    // 3. 获取 Nonce (防重放攻击的关键)
    // PendingNonceAt 返回该地址下一个可用的 nonce
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatal("Failed to get nonce:", err)
    }

    // 4. 估算 Gas Price 
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal("Failed to suggest gas price:", err)
    }

    // 5. 构建交易对象 (Raw Transaction)
    toAddress := common.HexToAddress(toStr)
    // Gas Limit 设置为 21000 (标准转账消耗)
    tx := types.NewTransaction(nonce, toAddress, amountWei, 21000, gasPrice, nil)

    // 6. 获取 ChainID (防止跨链重放攻击，比如把 ETH 的交易发到了 BSC 上)
    chainID, err := client.ChainID(context.Background())
    if err != nil {
        log.Fatal("Failed to get chainID:", err)
    }

    // 7. 离线签名 (这是最关键的一步！在本地签好名再发出去)
    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatal("Failed to sign tx:", err)
    }

    // 8. 广播交易
    // 注意：因为我们是测试，账号没钱，这里肯定会报错 "insufficient funds"
    // 但只要代码走到这一步，逻辑就是通的！
    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Printf("⚠️ Transaction failed (expected if no ETH): %v", err)
        return ""
    }

    log.Printf("✅ Transaction Sent! Hash: %s", signedTx.Hash().Hex())
    return signedTx.Hash().Hex()
}