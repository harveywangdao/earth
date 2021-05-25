package main

import (
	"flag"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	addr := flag.String("addr", "0.0.0.0:8545", "addr")
	flag.Parse()
	raddr, err := net.ResolveUDPAddr("udp4", *addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("udp server addr:", raddr)
	conn, err := net.DialUDP("udp4", nil, raddr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("ping"))
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("send ping")

	data := make([]byte, 128)
	n, _, err := conn.ReadFromUDP(data)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("recv:", string(data[:n]))
}
