package main

import (
	"log"
	"net"
	"time"
)

func server() {
	addr, err := net.ResolveUnixAddr("unix", "xiaoming")
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.ListenUnix("unix", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	log.Println("wait Accept")
	conn, err := lis.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	data := make([]byte, 128)
	log.Println("wait Read")
	log.Println(conn.Read(data))
	log.Println(string(data))

	time.Sleep(time.Second*20)

	data = []byte("i am server")
	log.Println(conn.Write(data))
}

func do1() {
	server()
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
