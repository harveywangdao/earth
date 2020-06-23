package main

import (
	"log"
	"time"

	"go.uber.org/ratelimit"
)

func do1() {
	rl := ratelimit.New(1) // per second

	prev := time.Now()
	for i := 0; i < 10; i++ {
		now := rl.Take()
		log.Println(i, now.Sub(prev))
		prev = now
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
