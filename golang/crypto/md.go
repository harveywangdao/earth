package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

func main() {
	s, _ := ioutil.ReadFile("20200401221752.jpg")
	mh := md5.Sum([]byte(s))
	fmt.Println(mh)

	hmh := hex.EncodeToString(mh[:])
	fmt.Println(hmh)

	bmh := base64.StdEncoding.EncodeToString(mh[:])
	fmt.Println(bmh)
	b, _ := base64.StdEncoding.DecodeString(bmh)
	fmt.Println(b)
}
