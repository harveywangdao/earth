package main

import (
	"log"

	"github.com/armon/go-socks5"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	if err := server.ListenAndServe("tcp", ":8080"); err != nil {
		panic(err)
	}
}
