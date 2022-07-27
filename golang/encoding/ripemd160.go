package main

import (
	"fmt"
	"golang.org/x/crypto/ripemd160"
)

func main() {
	data := "abcd1234"

	rip := ripemd160.New()
	rip.Write([]byte(data))

	hash := rip.Sum(nil)

	for i, v := range hash {
		fmt.Printf("%x ", v)
		if (i+1)%10 == 0 {
			fmt.Println()
		}
	}
}
