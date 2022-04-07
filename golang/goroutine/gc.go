package main

import (
	"log"
	"runtime"
	"time"
)

type Road int

func findRoad(r *Road) {
	log.Println("road:", *r)
}

func entry() {
	var rd Road = Road(999)
	r := &rd

	runtime.SetFinalizer(r, findRoad)
}

func do1() {
	entry()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		runtime.GC()
	}
}

/*
go tool trace --help
go tool compile --help
go tool link --help
go build -gcflags "-N -l" -o ggcc gc.go
GODEBUG="gctrace=1" ./ggcc
*/
type Data struct {
	d [1024 * 100]byte
	o *Data
}

func test() {
	var a, b Data
	a.o = &b
	b.o = &a
	runtime.SetFinalizer(&a, func(d *Data) { log.Printf("a %p final.\n", d) })
	runtime.SetFinalizer(&b, func(d *Data) { log.Printf("b %p final.\n", d) })
}

func do2() {
	for {
		test()
		time.Sleep(time.Millisecond)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do2()
}
