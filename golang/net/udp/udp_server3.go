package main

import (
	"log"
	"net"
)

func do1() {
	lis, err := net.ListenPacket("udp", ":9651")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()
	log.Println("udp server listen:", lis.LocalAddr())

	for {
		buf := make([]byte, 128)
		n, raddr, err := lis.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("udp server recv msg: %s", buf[:n])

		_, err = lis.WriteTo([]byte("hello, I am udp server"), raddr)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func do2() {
	laddr, err := net.ResolveUDPAddr("udp", ":9651")
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.ListenUDP("udp", laddr)
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()
	log.Println("udp server listen:", lis.LocalAddr())

	for {
		buf := make([]byte, 128)
		n, raddr, err := lis.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("udp server recv msg: %s", buf[:n])

		_, err = lis.WriteTo([]byte("hello, I am udp server"), raddr)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do2()
}
