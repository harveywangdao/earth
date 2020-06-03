package main

import (
	"fmt"
	"net/http"
	"time"
	//"runtime"

	_ "net/http/pprof"
)

func do() {
	data := make([]int, 1024)
	data[0] = 111
	time.Sleep(time.Hour)
}

func do2() {
	data := make([]int, 1024)
	data[0] = 222
	time.Sleep(time.Hour)
}

func do3() {
	data := make([]int, 1024)
	data[0] = 333
	time.Sleep(time.Hour)
}

func main() {
	go func() {
		count := 0
		for {
			go do()
			go do2()
			go do3()
			count++
			fmt.Println("count:", count)
			time.Sleep(time.Second)
		}
	}()

	http.ListenAndServe("0.0.0.0:6061", nil)
}
