package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	data := "abcd1234"

	sha := sha256.New()
	sha.Write([]byte(data))

	hash := sha.Sum(nil)

	for i, v := range hash {
		fmt.Printf("%x ", v)
		if (i+1)%8 == 0 {
			fmt.Println()
		}
	}
}
