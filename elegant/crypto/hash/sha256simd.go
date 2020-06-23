package main

import (
	"log"

	"github.com/minio/sha256-simd"
)

func do1() {
	s := sha256.New()
	s.Write([]byte("sadfsadf"))
	ha := s.Sum(nil)
	log.Println(ha)
}

func do2() {
	server := sha256.NewAvx512Server()
	h512 := sha256.NewAvx512(server)
	h512.Write([]byte("sadfsadf"))
	digest := h512.Sum([]byte{})
	log.Println(digest)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
