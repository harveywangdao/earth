package main

import (
	"log"
	"time"

	"github.com/dgraph-io/ristretto"
)

func do1() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		log.Fatal(err)
	}

	cache.Set("key", "valueww", 1)

	time.Sleep(10 * time.Millisecond)

	value, found := cache.Get("key")
	if !found {
		log.Println("missing value")
	}

	log.Println(value)
	cache.Del("key")
}

func do2() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		log.Fatal(err)
	}

	ok := cache.Set("key01", "value01", 1)
	log.Println(ok)
	ok = cache.Set("key02", "value02", 1)
	log.Println(ok)

	time.Sleep(10 * time.Millisecond)

	if value, found := cache.Get("key01"); found {
		log.Println(value)
	} else {
		log.Println("missing value")
	}

	if value, found := cache.Get("key02"); found {
		log.Println(value)
	} else {
		log.Println("missing value")
	}
}

func do3() {
	cache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})
	if err != nil {
		log.Fatal(err)
	}

	ok := cache.SetWithTTL("key01", "value01", 1, 5*time.Second)
	log.Println(ok)
	ok = cache.SetWithTTL("key02", "value02", 1, 5*time.Second)
	log.Println(ok)

	time.Sleep(10 * time.Millisecond)

	if value, found := cache.Get("key01"); found {
		log.Println(value)
	} else {
		log.Println("missing value")
	}

	if value, found := cache.Get("key02"); found {
		log.Println(value)
	} else {
		log.Println("missing value")
	}

	time.Sleep(6 * time.Second)

	if value, found := cache.Get("key01"); found {
		log.Println(value)
	} else {
		log.Println("missing value")
	}

	if value, found := cache.Get("key02"); found {
		log.Println(value)
	} else {
		log.Println("missing value")
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do3()
}
