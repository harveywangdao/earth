package main

import (
	"log"

	"github.com/hashicorp/golang-lru"
)

func do1() {
	c, err := lru.NewWithEvict(10, func(key interface{}, value interface{}) {
		log.Println("disapper", key, value)
	})
	//c, err := lru.New(10)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 20; i++ {
		c.Add(i, i*10)
	}

	log.Println(c.Len())

	for i := 0; i < 20; i++ {
		v, ok := c.Get(i)
		if ok {
			log.Println(v)
		}
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
