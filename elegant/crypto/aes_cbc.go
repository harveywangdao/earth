package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AESCBCEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize() // 16 bytes
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AESCBCDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize() // 16 bytes
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func main() {
	secretKey := []byte("1234567890123456")
	text := []byte("abc")

	crypted, err := AESCBCEncrypt(text, secretKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("aes128-cbc(abc)", hex.EncodeToString(crypted))
	fmt.Println("aes128-cbc(abc)", base64.StdEncoding.EncodeToString(crypted))

	plaintext, err := AESCBCDecrypt(crypted, secretKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("aes128-cbc plaintext", string(plaintext))
}
