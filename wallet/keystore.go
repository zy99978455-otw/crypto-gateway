package wallet

import (
    "github.com/ethereum/go-ethereum/accounts/keystore"
)

// CreateKeystore 生成一个新的加密钱包
// password: 你的提款密码
// dir: 密钥文件存放的文件夹路径
func CreateKeystore(password string, dir string) (string, error) {
    // 使用 StandardScryptN 和 P 参数
    // 这是为了防止黑客使用显卡(GPU)暴力破解密码
    ks := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)

    // 生成新账户
    account, err := ks.NewAccount(password)
    if err != nil {
        return "", err
    }

    // 返回生成的 0x 地址
    return account.Address.Hex(), nil
}