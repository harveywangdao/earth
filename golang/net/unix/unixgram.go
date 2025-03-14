package main

import (
	"log"
	"net"
	"os"
	"time"
)

const (
	sockFile = "/tmp/xyc.sock"
)

func server() {
	_ = os.Remove(sockFile)
	laddr, err := net.ResolveUnixAddr("unixgram", sockFile)
	if err != nil {
		log.Fatal(err)
		return
	}
	//net.ListenUnix() // unix/unixpacket
	//net.ListenPacket() // unixgram
	//net.Listen() // unix/unixpacket
	conn, err := net.ListenUnixgram("unixgram", laddr) // unixgram
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	data := make([]byte, 16)
	//conn.Read()
	//conn.ReadMsgUnix()
	n, addr, err := conn.ReadFrom(data)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("server recv: %s, addr: %v", data[:n], addr)
	/*	n, addr, err = conn.ReadFromUnix(data)
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println(n, addr, err)*/
	if _, err := conn.WriteTo([]byte("I am server"), addr); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 2)
}

func client() {
	time.Sleep(time.Second * 2)
	laddr, err := net.ResolveUnixAddr("unixgram", "/tmp/xya.sock")
	if err != nil {
		log.Fatal(err)
		return
	}
	raddr, err := net.ResolveUnixAddr("unixgram", sockFile)
	if err != nil {
		log.Fatal(err)
		return
	}

	// net.Dial("unixgram")
	conn, err := net.DialUnix("unixgram", laddr, raddr)
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = conn.Write([]byte("I am client"))
	if err != nil {
		log.Fatal(err)
		return
	}
	data := make([]byte, 16)
	n, addr, err := conn.ReadFrom(data)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("client recv: %s, addr: %v", data[:n], addr)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	go server()
	client()
}
