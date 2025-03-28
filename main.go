package main

import (
    "fmt"
    "log"
    
    // 引入钱包包
    "crypto-gateway/wallet" 
)

func main() {
    // 定义一个强密码
    password := "my_secret_password_123"
    
    // 定义存放目录 (记得在 gitignore 里忽略这个目录！)
    keyDir := "./tmp_keys"

    fmt.Println("Generating new keystore...")
    
    // 调用生成函数
    address, err := wallet.CreateKeystore(password, keyDir)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("✅ New Account Created: %s\n", address)
    fmt.Println("Key file saved in ./tmp_keys directory")
}