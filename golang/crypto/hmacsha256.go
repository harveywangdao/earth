package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	secretKey := []byte("dsde")
	text := []byte("abc|123|2020")
	h := hmac.New(sha256.New, secretKey)
	h.Write(text)
	signature := hex.EncodeToString(h.Sum(nil))

	fmt.Println("hmacsha256(abc|123|2020):", signature)
}
