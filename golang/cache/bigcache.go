package main

import (
	"log"
	"time"

	"github.com/allegro/bigcache"
)

func do1() {
	cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(5 * time.Second))
	if err != nil {
		log.Fatal(err)
		return
	}

	cache.Set("my-unique-key", []byte("value2"))

	entry, err := cache.Get("my-unique-key")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(entry))

	time.Sleep(10 * time.Second)

	entry, err = cache.Get("my-unique-key")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(entry))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
