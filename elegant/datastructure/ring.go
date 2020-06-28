package main

import (
	"container/ring"
	"fmt"
	"log"
)

func do1() {
	r := ring.New(10)

	for i := 0; i < 10; i++ {
		r.Value = i
		r = r.Next()
	}

	for i := 0; i < r.Len(); i++ {
		log.Println(r.Value)
		r = r.Next()
	}

	fmt.Println()

	r = r.Move(6)
	log.Println(r.Value)
	fmt.Println()

	for i := 0; i < r.Len(); i++ {
		log.Println(r.Value)
		r = r.Next()
	}
	fmt.Println()

	r1 := r.Unlink(11)
	for i := 0; i < r1.Len(); i++ {
		log.Println(r1.Value)
		r1 = r1.Next()
	}

	fmt.Println()

	for i := 0; i < r.Len(); i++ {
		log.Println(r.Value)
		r = r.Next()
	}

	fmt.Println()
	log.Println(r1.Len())
	log.Println(r.Len())
	fmt.Println()

	r.Do(func(i interface{}) {
		log.Println(i)
	})
	fmt.Println()

	r.Link(r1)

	r.Do(func(i interface{}) {
		log.Println(i)
	})
	fmt.Println()

	r1.Do(func(i interface{}) {
		log.Println(i)
	})
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
