package main

import (
	"log"
	"net"
)

func udpDail() {
	laddr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:10002")
	if err != nil {
		log.Fatal(err)
		return
	}
	conn, err := net.ListenUDP("udp4", laddr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	raddr, err := net.ResolveUDPAddr("udp4", "47.239.67.241:10001")
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = conn.WriteToUDP([]byte("udp client say hello"), raddr)
	if err != nil {
		log.Fatal(err)
		return
	}

	buf := make([]byte, 1024)
	n, raddr2, err := conn.ReadFrom(buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(raddr2)
	log.Printf("public ip: %s", buf[:n])
	n, raddr2, err = conn.ReadFrom(buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(raddr2)
	log.Printf("udp client recv: %s", buf[:n])
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	udpDail()
}
