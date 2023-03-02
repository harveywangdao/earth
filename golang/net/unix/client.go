package main

import (
	"log"
	"net"
)

func client() {
	addr, err := net.ResolveUnixAddr("unix", "xiaoming")
	if err != nil {
		log.Fatal(err)
	}
	laddr, err := net.ResolveUnixAddr("unix", "xiaoming_client")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := net.DialUnix("unix", laddr, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	data := []byte("i am client")
	log.Println(conn.Write(data))
	//conn.WriteTo(b, addr)
	//conn.WriteToUnix(b, addr)
	//conn.WriteMsgUnix(b, oob, addr)

	data = make([]byte, 128)
	log.Println(conn.Read(data))
	log.Println(string(data))
}

func do1() {
	client()
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
