package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"hash/crc32"
)

func main() {
	m4 := md4.New()
	m4.Write([]byte("This is md4"))
	fmt.Println("md4 =", m4.Sum(nil))

	m5 := md5.New()
	m5.Write([]byte("This is md5"))
	fmt.Println("md5 =", m5.Sum(nil))

	s256 := sha256.New()
	s256.Write([]byte("This is sha256"))
	fmt.Println("sha256 =", s256.Sum(nil))

	r160 := ripemd160.New()
	r160.Write([]byte("This is ripemd160"))
	fmt.Println("ripemd160 =", r160.Sum(nil))

	c32 := crc32.NewIEEE()
	c32.Write([]byte("This is crc32"))
	fmt.Println("crc32 =", c32.Sum(nil))
	//fmt.Printf("crc32 = 0x%X\n", crc32.ChecksumIEEE([]byte("This is crc32")))

	h := hmac.New(sha256.New, []byte("123456"))
	h.Write([]byte("This is hmac sha256"))
	fmt.Println("hmac sha256 =", h.Sum(nil))
}
