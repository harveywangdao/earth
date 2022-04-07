package main

import (
	"log"

	"github.com/dgrijalva/lfu-go"
)

func do1() {
	c := lfu.New()
	c.Set("key1", 1)
	c.Set("key1", 1)
	c.Set("key1", 1)

	c.Set("key2", 1)
	c.Set("key2", 1)

	c.Evict(1)

	log.Println(c.Get("key1"))
	log.Println(c.Get("key2"))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
