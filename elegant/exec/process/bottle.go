package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("sleep...")
	time.Sleep(time.Second * 5)
	fmt.Println("sleep stop...")
	var p *int
	*p = 2
}
