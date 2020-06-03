package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	do1()
	return

	s := "AAAAAAAAA"
	bs := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println("after encoding:", bs)
	ds, err := base64.StdEncoding.DecodeString(bs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("after decoding:", string(ds))
}

func do1() {
	s := "fdsvf"

	d, _ := base64.StdEncoding.DecodeString(s)
	fmt.Println(len(d))
	fmt.Println(d)
}
