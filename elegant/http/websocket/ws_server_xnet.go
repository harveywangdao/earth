package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/websocket"
)

func upper(ws *websocket.Conn) {
	var err error
	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			log.Println(err)
			continue
		}

		if err = websocket.Message.Send(ws, strings.ToUpper(reply)); err != nil {
			log.Println(err)
			continue
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		return
	}

	t, _ := template.ParseFiles("ws_client_xnet.html")
	t.Execute(w, nil)
}

func do1() {
	http.Handle("/upper", websocket.Handler(upper))
	http.HandleFunc("/", index)

	if err := http.ListenAndServe(":8544", nil); err != nil {
		log.Fatal(err)
	}
}

func echoServer(ws *websocket.Conn) {
	go func() {
		buf := make([]byte, 100)
		for {
			n, err := ws.Read(buf)
			if err != nil {
				log.Println(err)
				break
			}
			log.Println("server receive: ", string(buf[:n]))
		}
	}()
	var send int
	for {
		sendStr := strconv.Itoa(send)
		_, err := ws.Write([]byte("I am server " + sendStr))
		if err != nil {
			log.Println(err)
			break
		}
		log.Println("server send: ", sendStr)
		time.Sleep(time.Second)
		send++
	}
}

func do2() {
	http.Handle("/echo", websocket.Handler(echoServer))

	if err := http.ListenAndServe(":8544", nil); err != nil {
		log.Fatal(err)
	}
}

func Copy(ws *websocket.Conn) {
	log.Printf("copyServer %#v\n", ws.Config())
	io.Copy(ws, ws)
	log.Println("copyServer finished")
}

func ReadWrite(ws *websocket.Conn) {
	log.Printf("readWriteServer %#v\n", ws.Config())
	buf := make([]byte, 100)
	for {
		n, err := ws.Read(buf)
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("recv:%q\n", buf[:n])
		n, err = ws.Write(buf[:n])
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("send:%q\n", buf[:n])
	}
	log.Println("readWriteServer finished")
}

func RecvSend(ws *websocket.Conn) {
	log.Printf("recvSendServer %#v\n", ws)
	for {
		var buf string
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("recv:%q\n", buf)
		err = websocket.Message.Send(ws, buf)
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("send:%q\n", buf)
	}
	log.Println("recvSendServer finished")
}

func RecvSendBinary(ws *websocket.Conn) {
	log.Printf("recvSendBinaryServer %#v\n", ws)
	for {
		var buf []byte
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("recv:%#v\n", buf)
		err = websocket.Message.Send(ws, buf)
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("send:%#v\n", buf)
	}
	log.Println("recvSendBinaryServer finished")
}

func Json(ws *websocket.Conn) {
	type T struct {
		Msg  string `json:"msg"`
		Path string `json:"path"`
	}

	log.Printf("jsonServer %#v\n", ws.Config())
	for {
		var msg T
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("recv:%#v\n", msg)
		err = websocket.JSON.Send(ws, msg)
		if err != nil {
			log.Println(err)
			break
		}
		log.Printf("send:%#v\n", msg)
	}
	log.Println("jsonServer finished")
}

func web(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("ws_client_xnet2.html")
	t.Execute(w, nil)
}

func do3() {
	http.Handle("/copy", websocket.Handler(Copy))
	http.Handle("/readWrite", websocket.Handler(ReadWrite))
	http.Handle("/recvSend", websocket.Handler(RecvSend))
	http.Handle("/recvSendBinary", websocket.Handler(RecvSendBinary))
	http.Handle("/json", websocket.Handler(Json))

	http.HandleFunc("/web", web)

	if err := http.ListenAndServe(":8544", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do3()
}
