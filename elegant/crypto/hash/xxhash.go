package main

import (
	"log"

	"github.com/cespare/xxhash"
)

func do1() {
	ha := xxhash.Sum64([]byte("adsadas"))
	log.Println(ha)

	ha = xxhash.Sum64String("adsadas")
	log.Println(ha)

	d := xxhash.New()
	d.Write([]byte("adsadas"))
	log.Println(d.Sum(nil))
	log.Println(d.Sum64())
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
