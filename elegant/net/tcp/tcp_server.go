package main

import (
	"log"
	"net"
	"time"

	"github.com/felixge/tcpkeepalive"
)

func do1() {
	lis, err := net.Listen("tcp", ":8564")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()
	log.Println("tcp server listen:", lis.Addr())

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()

			buf := make([]byte, 128)
			n, err := c.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("tcp server recv msg: %s", buf[:n])

			_, err = c.Write([]byte("hello, I am tcp server"))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Hour)
		}(conn)
	}
}

func do2() {
	lis, err := net.Listen("tcp", ":8564")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()
	log.Println("tcp server listen:", lis.Addr())

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()

			tcp, ok := c.(*net.TCPConn)
			if !ok {
				log.Fatal("not tcp")
			}
			if err := tcp.SetKeepAlive(true); err != nil {
				log.Fatal(err)
			}
			if err := tcp.SetKeepAlivePeriod(7 * time.Second); err != nil {
				log.Fatal(err)
			}

			buf := make([]byte, 128)
			n, err := c.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("tcp server recv msg: %s", buf[:n])

			_, err = c.Write([]byte("hello, I am tcp server"))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Hour)
		}(conn)
	}
}

func do3() {
	lis, err := net.Listen("tcp", ":8564")
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()
	log.Println("tcp server listen:", lis.Addr())

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()

			newConn, err := tcpkeepalive.EnableKeepAlive(c)
			if err != nil {
				log.Fatal("EnableKeepAlive failed:", err)
			}
			err = newConn.SetKeepAliveIdle(11 * time.Second)
			if err != nil {
				log.Fatal("SetKeepAliveIdle failed:", err)
			}
			err = newConn.SetKeepAliveCount(3)
			if err != nil {
				log.Fatal("SetKeepAliveCount failed:", err)
			}
			err = newConn.SetKeepAliveInterval(7 * time.Second)
			if err != nil {
				log.Fatal("SetKeepAliveInterval failed:", err)
			}

			buf := make([]byte, 128)
			n, err := newConn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("tcp server recv msg: %s", buf[:n])

			_, err = newConn.Write([]byte("hello, I am tcp server"))
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Hour)
		}(conn)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
