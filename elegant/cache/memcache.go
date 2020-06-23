package main

import (
	"log"
	"strconv"
	"sync"

	"github.com/bradfitz/gomemcache/memcache"
)

func do1() {
	mc := memcache.New("192.168.197.128:11211")
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	it, err := mc.Get("foo")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(it.Key, string(it.Value))

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			it.Value = []byte("addcalskc" + strconv.Itoa(i))
			err := mc.CompareAndSwap(it)
			if err != nil {
				log.Println(err)
				return
			}
		}(i)
	}

	wg.Wait()

	it, err = mc.Get("foo")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(it.Key, string(it.Value))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
