package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/google/btree"
)

func do1() {
	vals := rand.Perm(10000000)
	tree := btree.New(8)
	start := time.Now()
	for _, v := range vals {
		tree.ReplaceOrInsert(btree.Int(v))
	}
	log.Println("insert btree cost", time.Since(start))

	value := 3333
	start = time.Now()
	log.Println(tree.Get(btree.Int(value)))
	log.Println("btree get cost", time.Since(start))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
