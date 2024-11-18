package main

import (
	"log"
	"net"
)

func tcpServ() {
	listener, err := net.Listen("tcp4", ":9298")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer listener.Close()
	log.Printf("tcp server listen: %v", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println(conn.LocalAddr(), conn.RemoteAddr())
		_, err = conn.Write([]byte(conn.RemoteAddr().String()))
		if err != nil {
			log.Println(err)
		}
		conn.Close()
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	tcpServ()
}
