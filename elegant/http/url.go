package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://www.baidu.com/p?a=1&b=3")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("EscapedPath =", u.EscapedPath())
	fmt.Println("Hostname =", u.Hostname())
	fmt.Println("Port =", u.Port())
	fmt.Println("Query =", u.Query())
	fmt.Println("Host =", u.Host)
	fmt.Println("Opaque =", u.Opaque)
	fmt.Println("Path =", u.Path)
	fmt.Println("RawPath =", u.RawPath)
	fmt.Println("RawQuery =", u.RawQuery)
	fmt.Println("Scheme =", u.Scheme)
	fmt.Println("User =", u.User)
}
