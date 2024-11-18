package main

import (
	"log"
	"net"
)

func tcpDail() {
	laddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:9297")
	if err != nil {
		log.Fatal(err)
		return
	}
	raddr, err := net.ResolveTCPAddr("tcp4", "47.239.67.241:9298")
	if err != nil {
		log.Fatal(err)
		return
	}
	conn, err := net.DialTCP("tcp4", laddr, raddr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println(conn.LocalAddr(), conn.RemoteAddr())
	log.Printf("public ip: %s", buf[:n])
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	tcpDail()
}
