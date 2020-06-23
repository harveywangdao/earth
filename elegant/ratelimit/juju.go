package main

import (
	"log"
	"sync"
	"time"

	jujuratelimit "github.com/juju/ratelimit"
)

func do1() {
	bucket := jujuratelimit.NewBucket(time.Second, 10)
	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			log.Println("goroutine", i, "waiting")
			bucket.Wait(1)

			log.Println("goroutine", i, "get token")
		}(i)
	}

	wg.Wait()
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
