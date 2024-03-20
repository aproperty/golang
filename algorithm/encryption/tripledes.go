package encryption

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

// 加密
func TripleDesEncrypt(orig string, key string) string {

	// 3 DES 的秘钥长度必须为24位
	k := []byte(key)
	block, _ := des.NewTripleDESCipher(k)

	origData := []byte(orig)
	origData = PKCS7Padding(origData, block.BlockSize()) // 补全码

	crypted := make([]byte, len(origData)) // 创建密文数组

	blockMode := cipher.NewCBCEncrypter(block, k[:8]) // 设置加密方式
	blockMode.CryptBlocks(crypted, origData)          // 加密

	return base64.StdEncoding.EncodeToString(crypted)
}

// 解密
func TipleDesDecrypt(crypted string, key string) string {

	k := []byte(key)
	block, _ := des.NewTripleDESCipher(k)

	blockMode := cipher.NewCBCDecrypter(block, k[:8])

	cryptedByte, _ := base64.StdEncoding.DecodeString(crypted)
	origData := make([]byte, len(cryptedByte))
	blockMode.CryptBlocks(origData, cryptedByte)

	origData = PKCS7UnPadding(origData)

	return string(origData)
}
