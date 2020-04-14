package zz

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

//aes 解密
func AESDecrypt(crypted, key []byte, iv []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv[:blockSize])
	orgData := make([]byte, len(crypted))
	blockMode.CryptBlocks(orgData, crypted)
	orgData = PKCS7UnPadding(orgData)

	return string(orgData), nil
}

//pkcs7 去掉补码
func PKCS7UnPadding(orgData []byte) []byte {
	length := len(orgData)
	unpadding := int(orgData[length-1])
	return orgData[:length-unpadding]
}

//加密
func AESEncrypt(orgData, key []byte, iv []byte) []byte {
	//获取block块
	block, _ := aes.NewCipher(key)
	//补码
	orgData = PKCS7Padding(orgData, block.BlockSize())
	//加密模式，
	blockMode := cipher.NewCBCEncrypter(block, iv[:block.BlockSize()])
	//创建明文长度的数组
	crypted := make([]byte, len(orgData))
	//加密明文
	blockMode.CryptBlocks(crypted, orgData)
	return crypted
}

//pkcs7 补码
func PKCS7Padding(orgData []byte, blockSize int) []byte {
	//计算需要补几位数
	padding := blockSize - len(orgData)%blockSize
	//在切片后面追加char数量的byte(char)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(orgData, padtext...)
}
