package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listenner, err := net.Listen("tcp", ":7849")
	if err != nil {
		log.Fatal(err)
	}
	defer listenner.Close()

	for {
		conn, err := listenner.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		go func(c net.Conn) {
			defer c.Close()
			ipAddr := c.RemoteAddr().String()
			fmt.Println(ipAddr, "连接成功")

			buf := make([]byte, 4096)
			n, err := c.Read(buf)
			if err != nil {
				fmt.Println(err)
				return
			}

			result := buf[:n]
			fmt.Printf("接收到数据来自[%s]==>:\n%s\n", ipAddr, string(result))
		}(conn)
	}
}
