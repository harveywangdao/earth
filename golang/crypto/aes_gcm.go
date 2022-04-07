package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func AESGCMEncrypt(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aesgcm.NonceSize())
	rand.Read(nonce)

	fmt.Println("aesgcm.NonceSize():", aesgcm.NonceSize())

	return aesgcm.Seal(nonce, nonce, plaintext, nil), nil
}

func AESGCMDecrypt(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return aesgcm.Open(nil, ciphertext[:aesgcm.NonceSize()], ciphertext[aesgcm.NonceSize():], nil)
}

func main() {
	secretKey := []byte("1234567890123456")
	text := []byte("abc")

	crypted, err := AESGCMEncrypt(text, secretKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("aes128-gcm(abc)", hex.EncodeToString(crypted))
	fmt.Println("aes128-gcm(abc)", base64.StdEncoding.EncodeToString(crypted))

	plaintext, err := AESGCMDecrypt(crypted, secretKey)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("aes128-gcm plaintext", string(plaintext))
}
