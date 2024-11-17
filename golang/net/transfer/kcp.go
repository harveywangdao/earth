package main

import (
	"crypto/rand"
	"fmt"
	"github.com/xtaci/kcp-go/v5"
	"github.com/xtaci/kcptun/std"
	"log"
	"net"
	"time"
)

const (
	tcpServerPort = ":8080"
	tcpClientPort = ":8388"
)

func size(n int) string {
	if n >= 1024*1024*1024 {
		return fmt.Sprintf("%dGB", n/1024/1024/1024)
	} else if n >= 1024*1024 {
		return fmt.Sprintf("%dMB", n/1024/1024)
	} else if n >= 1024 {
		return fmt.Sprintf("%dKB", n/1024)
	}
	return fmt.Sprintf("%dByte", n)
}

func tcpClient1() {
	conn, err := net.Dial("tcp", "127.0.0.1"+tcpClientPort)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	_, err = conn.Write([]byte("hello, I am tcp client"))
	if err != nil {
		log.Fatal(err)
		return
	}
	buf := make([]byte, 128)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("tcp client recv msg: %s", buf[:n])
}

func tcpClient() {
	conn, err := net.Dial("tcp", "127.0.0.1"+tcpClientPort)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 1024*1024)
	var total int
	start := time.Now()
	for {
		n, err := rand.Read(buf)
		if err != nil {
			log.Fatal(err)
			return
		}
		if _, err := conn.Write(buf[:n]); err != nil {
			log.Fatal(err)
			return
		}
		total += n
		cost := time.Now().Sub(start).Seconds()
		speed := float64(total) / cost
		log.Printf("tcp client, send size: %d, total size: %s, cost: %v, speed: %fMB/s", n, size(total), cost, speed/1024/1024)
	}
}

func tcpServer() {
	lis, err := net.Listen("tcp", tcpServerPort)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer lis.Close()
	log.Println("tcp server listen:", lis.Addr())
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		go tcpHandle(conn)
	}
}

func tcpHandle(c net.Conn) {
	defer c.Close()
	buf := make([]byte, 1024*1024)
	var total int
	start := time.Now()
	for {
		n, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		total += n
		cost := time.Now().Sub(start).Seconds()
		speed := float64(total) / cost
		log.Printf("tcp server, recv size: %d, total size: %s, cost: %v, speed: %fMB/s", n, size(total), cost, speed/1024/1024)
	}
}

func tcpTest() {
	go tcpServer()
	time.Sleep(time.Millisecond)
	tcpClient()
}

func udpServer() {
	addr, err := net.ResolveUDPAddr("udp4", ":9090")
	if err != nil {
		log.Fatal(err)
		return
	}
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	log.Println("udp server start, listen:", conn.LocalAddr())
	buf := make([]byte, 1024*32)
	var total int
	start := time.Now()
	for {
		n, raddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		total += n
		cost := time.Now().Sub(start).Seconds()
		speed := float64(total) / cost
		log.Printf("udp server, recv size: %d, total size: %s, cost: %v, speed: %fMB/s", n, size(total), cost, speed/1024/1024)

		ack := []byte("ack")
		_, err = conn.WriteToUDP(ack, raddr)
		if err != nil {
			log.Println(err)
		}
		//conn.WriteToUDP()
		//conn.WriteTo()
		//conn.WriteMsgUDP()
		//conn.WriteMsgUDPAddrPort()
		//conn.WriteToUDPAddrPort()
		//conn.Write
	}
}

func udpClient() {
	addr, err := net.ResolveUDPAddr("udp4", "127.0.0.1:9090")
	if err != nil {
		log.Fatal(err)
		return
	}
	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024*32)
	var total int
	start := time.Now()
	for {
		n, err := rand.Read(buf)
		if err != nil {
			log.Fatal(err)
			return
		}
		if _, err := conn.Write(buf[:n]); err != nil {
			log.Fatal(err)
			return
		}
		total += n
		cost := time.Now().Sub(start).Seconds()
		speed := float64(total) / cost
		log.Printf("udp client, send size: %d, total size: %s, cost: %v, speed: %fMB/s", n, size(total), cost, speed/1024/1024)

		n, err = conn.Read(buf)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("udp client, recv: %s", buf[:n])
		}
		//conn.Read()
		//conn.ReadFrom()
		//conn.ReadFromUDP()
		//conn.ReadMsgUDP()
		//conn.ReadFromUDPAddrPort()
		//conn.ReadMsgUDPAddrPort()
	}
}

func udpServer1() {
	conn, err := net.ListenPacket("udp4", ":9090")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	log.Println("udp server start, listen:", conn.LocalAddr())
	buf := make([]byte, 1024*32)
	var total int
	start := time.Now()
	for {
		n, addr, err := conn.ReadFrom(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println(addr)
		total += n
		cost := time.Now().Sub(start).Seconds()
		speed := float64(total) / cost
		log.Printf("udp server, recv size: %d, total size: %s, cost: %v, speed: %fMB/s", n, size(total), cost, speed/1024/1024)
	}
}

func udpClient1() {
	conn, err := net.Dial("udp4", "127.0.0.1:9090")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()
	buf := make([]byte, 1024*32)
	var total int
	start := time.Now()
	for {
		n, err := rand.Read(buf)
		if err != nil {
			log.Fatal(err)
			return
		}
		if _, err := conn.Write(buf[:n]); err != nil {
			log.Fatal(err)
			return
		}
		total += n
		cost := time.Now().Sub(start).Seconds()
		speed := float64(total) / cost
		log.Printf("udp client, send size: %d, total size: %s, cost: %v, speed: %fMB/s", n, size(total), cost, speed/1024/1024)
	}
}

func udpTest() {
	go udpServer1()
	time.Sleep(time.Millisecond)
	udpClient1()
}

func kcpServer() {
	//key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	//block, _ := kcp.NewAESBlockCrypt(key)
	listener, err := kcp.ListenWithOptions("127.0.0.1:12345", nil, 10, 3)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer listener.Close()
	for {
		s, err := listener.AcceptKCP()
		if err != nil {
			log.Fatal(err)
		}
		go kcpHandle(s)
	}
}

func kcpHandle(conn *kcp.UDPSession) {
	defer conn.Close()
	buf := make([]byte, 1024)
	var total int
	start := time.Now()
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
			return
		}
		total += n
		cost := time.Now().Sub(start).Seconds()
		speed := float64(total) / cost
		log.Printf("kcp server, recv size: %d, total size: %s, cost: %v, speed: %fMB/s", n, size(total), cost, speed/1024/1024)
	}
}

func kcpClient() {
	//key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	//block, _ := kcp.NewAESBlockCrypt(key)
	sess, err := kcp.DialWithOptions("127.0.0.1:12345", nil, 10, 3)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer sess.Close()
	buf := make([]byte, 1024)
	var total int
	start := time.Now()
	for {
		n, err := rand.Read(buf)
		if err != nil {
			log.Fatal(err)
			return
		}
		if _, err := sess.Write(buf[:n]); err != nil {
			log.Fatal(err)
			return
		}
		total += n
		cost := time.Now().Sub(start).Seconds()
		speed := float64(total) / cost
		log.Printf("kcp client, send size: %d, total size: %s, cost: %v, speed: %fMB/s", n, size(total), cost, speed/1024/1024)
	}
}

func kcpTest() {
	go kcpServer()
	time.Sleep(time.Millisecond)
	kcpClient()
}

func kcptunTest() {
	std.Pipe(nil, nil, 0)
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//udpTest()
	//tcpTest()
	kcpTest()
}
