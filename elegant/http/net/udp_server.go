package main

import (
	"flag"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	port := flag.String("port", "8545", "port")
	flag.Parse()

	addr, err := net.ResolveUDPAddr("udp4", ":"+*port)
	if err != nil {
		log.Fatal(err)
		return
	}
	lis, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer lis.Close()

	log.Println("udp server start, listen:", *port)
	for {
		data := make([]byte, 128)
		n, raddr, err := lis.ReadFromUDP(data)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("client addr:", raddr)
		log.Println("recv msg:", string(data[:n]))

		_, err = lis.WriteToUDP([]byte("pong"), raddr)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
