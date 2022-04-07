package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
)

func AesOfbDecrypt(key []byte, encryptedBytes, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ofb := cipher.NewOFB(block, iv)
	plainText := make([]byte, len(encryptedBytes))
	ofb.XORKeyStream(plainText, encryptedBytes)
	return plainText, nil
}

func AesOfbEncrypt(key, plainText []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := make([]byte, block.BlockSize())
	rand.Read(iv)
	ofb := cipher.NewOFB(block, iv)
	cipherText := make([]byte, len(plainText))
	ofb.XORKeyStream(cipherText, plainText)
	return cipherText, nil
}

func main() {
	secretKey := []byte("1234567890123456")
	text := []byte("abc")

	b, err := AesOfbEncrypt(secretKey, text)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("aes128-ofb(abc)", hex.EncodeToString(b))
	fmt.Println("aes128-ofb(abc)", base64.StdEncoding.EncodeToString(b))
}
