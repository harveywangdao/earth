package main

import (
	"log"
	"time"

	"github.com/patrickmn/go-cache"
)

func do1() {
	c := cache.New(5*time.Minute, 10*time.Minute)

	c.Set("foo", "bar", cache.DefaultExpiration)
	c.Set("baz", 42, cache.NoExpiration)

	foo, found := c.Get("foo")
	if found {
		log.Println(foo)
	}
}

func do2() {
	c := cache.New(5*time.Second, 10*time.Second)

	c.Set("foo", "bar", cache.DefaultExpiration)
	foo, found := c.Get("foo")
	if found {
		log.Println(foo)
	}

	time.Sleep(6 * time.Second)
	key, ok := c.Get("foo")
	if ok {
		log.Println(key)
	} else {
		log.Println("not existed")
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do2()
}
