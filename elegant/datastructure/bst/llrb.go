package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/petar/GoLLRB/llrb"
)

func do1() {
	vals := rand.Perm(1000000)
	tr := llrb.New()

	start := time.Now()
	for _, v := range vals {
		tr.ReplaceOrInsert(llrb.Int(v))
	}
	log.Println("insert cost:", time.Since(start))

	value := 500000

	start = time.Now()
	tr.Get(llrb.Int(value))
	log.Println("llrb get cost:", time.Since(start))

	start = time.Now()
	for _, v := range vals {
		if value == v {
			break
		}
	}
	log.Println("array get cost:", time.Since(start))
}

func do2() {

}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do2()
}
