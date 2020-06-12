package main

import (
	"log"
	"strconv"
	"time"

	"golang.org/x/net/websocket"
)

func do1() {
	origin := "http://localhost/"              //客户端地址
	url := "ws://localhost:8544/echo"          //服务器地址
	ws, err := websocket.Dial(url, "", origin) //第二个参数是websocket子协议，可以为空
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		var send int
		for {
			sendStr := strconv.Itoa(send)
			_, err := ws.Write([]byte("I am client " + sendStr))
			if err != nil {
				log.Fatal(err)
			}
			log.Println("client send: ", sendStr)
			time.Sleep(time.Second * 2)
			send++
		}
	}()

	var buf = make([]byte, 100)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("client receive: ", string(buf[:n]))
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
