package main

import (
	"container/list"
	"log"
)

func do1() {
	li := list.New()

	for i := 0; i < 10; i++ {
		li.PushBack(i)
	}

	for p := li.Front(); p != li.Back(); p = p.Next() {
		log.Println(p.Value)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
