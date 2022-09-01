package main

import (
	"log"
	"net"
	"sync"
	"time"
)

func client(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	addr, err := net.ResolveUnixAddr("unix", "xiaoming")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUnix("unix", nil, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	data := []byte("i am hello")
	log.Println(conn.Write(data))
	//conn.WriteTo(b, addr)
	//conn.WriteToUnix(b, addr)
	//conn.WriteMsgUnix(b, oob, addr)

	data = make([]byte, 128)
	log.Println(conn.Read(data))
	log.Println(string(data))
}

func server(wg *sync.WaitGroup) {
	defer wg.Done()
	addr, err := net.ResolveUnixAddr("unix", "xiaoming")
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.ListenUnix("unix", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	conn, err := lis.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	data := make([]byte, 128)
	log.Println(conn.Read(data))
	log.Println(string(data))

	data = []byte("i am world")
	log.Println(conn.Write(data))
}

func do1() {
	var wg sync.WaitGroup
	wg.Add(2)
	go server(&wg)
	go client(&wg)
	wg.Wait()
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
