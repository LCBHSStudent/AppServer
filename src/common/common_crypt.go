package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
)

// for cmd:GO TEST - unit test
func Crypt() error {
	origin := "chikubi"
	key    := "0123456789123456"
	fmt.Println("原文：", origin)
	
	encryptCode := AesEncrypt(origin, key)
	fmt.Println("密文：", encryptCode)
	
	decryptCode := AesDecrypt(encryptCode, key)
	fmt.Println(decryptCode)
	
	return nil
}

func AesEncrypt(origin string, key string) string {
	originData := []byte(origin)
	keyData    := []byte(key)
	// 分组密钥
	// NewCipher限制了keyData的长度必须为16，24或32
	block, err   := aes.NewCipher(keyData)
	if err != nil {
		panic(err)
	}
	// 获取密钥块的长度
	blockSize  := block.BlockSize()
	// 补全码
	originData = PKCS7Padding(originData, blockSize)
	// 加密模式
	blockMode  := cipher.NewCBCEncrypter(block, keyData[:blockSize])
	// 创建数组
	crypted     := make([]byte, len(originData))
	// 加密
	blockMode.CryptBlocks(crypted, originData)
	return base64.StdEncoding.EncodeToString(crypted)
}

func AesDecrypt(crypted string, key string) string {
	// base64转换为字节数组
	cryptedData, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		log.Fatal(err)
	}
	keyData        := []byte(key)
	// 分组密钥
	block, err       := aes.NewCipher(keyData)
	if err != nil {
		log.Fatal(err)
	}
	blockSize      := block.BlockSize()
	// 加密模式
	blockMode      := cipher.NewCBCDecrypter(block, keyData[:blockSize])
	// 创建数组
	originData     := make([]byte, len(cryptedData))
	// 解密
	blockMode.CryptBlocks(originData, cryptedData)
	// 去除补全码
	originData     = PKCS7UnPadding(originData)
	return string(originData)
}

// 加补全码
// AES加密数据块分组长度必须为128bit(byte[16]), 密钥长度可以是128bit(byte[16])
// 192bit(byte[24]), 256bit(byte[32])中的任意一个
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

// 去除补全码
func PKCS7UnPadding(originData[] byte) []byte {
	length    := len(originData)
	unpadding := int(originData[length - 1])
	return originData[:(length - unpadding)]
}