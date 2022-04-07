package main

import (
	"log"
	"net"
)

func do1() {
	conn, err := net.Dial("udp", "127.0.0.1:9651")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello, I am udp client"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("udp client recv msg: %s", buf[:n])
}

func do2() {
	raddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9651")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialUDP("udp", nil, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello, I am udp client"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("udp client recv msg: %s", buf[:n])
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do2()
}
