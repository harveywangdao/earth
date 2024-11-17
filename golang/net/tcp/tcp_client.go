package main

import (
	"log"
	"net"
	"time"
)

func c1() {
	conn, err := net.Dial("tcp", "127.0.0.1:8564")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello, I am tcp client"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("tcp client recv msg: %s", buf[:n])

	//time.Sleep(time.Second * 5)
	time.Sleep(time.Hour)
}

func c2() {
	var d net.Dialer
	d.KeepAlive = 3 * time.Second

	conn, err := d.Dial("tcp", "127.0.0.1:8564")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello, I am tcp client"))
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("tcp client recv msg: %s", buf[:n])

	//time.Sleep(time.Second * 5)
	time.Sleep(time.Hour)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	c1()
}
