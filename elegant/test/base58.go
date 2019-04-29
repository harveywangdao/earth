package main

import (
	"fmt"
	"github.com/mr-tron/base58/base58"
)

func main() {
	data := "12345678"
	encoded := base58.Encode([]byte(data))
	fmt.Println(encoded)

	decoded, err := base58.Decode(encoded)
	if err != nil {
		fmt.Println("2 :", err.Error())
		return
	}
	fmt.Println(string(decoded))
}
