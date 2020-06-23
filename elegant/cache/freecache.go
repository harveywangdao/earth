package main

import (
	"log"

	"github.com/coocood/freecache"
)

func do1() {
	cacheSize := 100 * 1024 * 1024
	cache := freecache.NewCache(cacheSize)
	//debug.SetGCPercent(20)

	key := []byte("abc")
	val := []byte("def")

	cache.Set(key, val, 60)
	got, err := cache.Get(key)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(string(got))
	}
	affected := cache.Del(key)
	log.Println("deleted key ", affected)
	log.Println("entry count ", cache.EntryCount())

	cache.Set(key, []byte("1111"), 60)
	got, err = cache.Get(key)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(string(got))
	}
	cache.Set(key, []byte("2222"), 60)
	got, err = cache.Get(key)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(string(got))
	}

	got, err = cache.Get([]byte("key"))
	if err != nil {
		log.Println(err)
	} else {
		log.Println(string(got))
	}

	log.Println(cache.Del([]byte("key")))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
