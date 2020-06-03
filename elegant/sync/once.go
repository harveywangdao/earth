package main

import (
	"fmt"
	"sync"
	"time"
)

func oncefunc() {
	fmt.Println("oncefunc doing something...")
}

func drink(o *sync.Once) {
	fmt.Println("Start drink")

	o.Do(oncefunc)

	fmt.Println("Drink end")
}

func main() {
	o := &sync.Once{}

	go drink(o)
	go drink(o)

	time.Sleep(time.Second * 2)
	drink(o)
	oncefunc()
}
