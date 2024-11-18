package main

import (
	"log"
	"net"
)

func updServ() {
	laddr, err := net.ResolveUDPAddr("udp4", ":10001")
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
	log.Printf("udp server listen: %v", conn.LocalAddr())
	buf := make([]byte, 1024)
	for {
		n, raddr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Printf("udp server recv msg: %s", buf[:n])
		_, err = conn.WriteTo([]byte(raddr.String()), raddr)
		if err != nil {
			log.Println(err)
		}
		udpServerToClient(raddr)
	}
}

func udpServerToClient(addr net.Addr) {
	laddr, err := net.ResolveUDPAddr("udp4", "0.0.0.0:10003")
	if err != nil {
		log.Fatal(err)
		return
	}
	raddr, ok := addr.(*net.UDPAddr)
	if !ok {
		log.Fatal(addr)
		return
	}
	conn, err := net.DialUDP("udp4", laddr, raddr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	log.Printf("udp server to client, %v -> %v", conn.LocalAddr(), conn.RemoteAddr())
	_, err = conn.Write([]byte("udp server say hello"))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	updServ()
}
