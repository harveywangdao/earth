package main

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"

	quic "github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

func do1() {
	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	session, err := quic.DialAddr("localhost:8586", tlsConf, nil)
	if err != nil {
		log.Println(err)
		return
	}

	stream, err := session.OpenStreamSync(context.Background())
	if err != nil {
		log.Println(err)
		return
	}

	message := "xiaoming"
	log.Printf("Client: Sending '%s'\n", message)
	_, err = stream.Write([]byte(message))
	if err != nil {
		log.Println(err)
		return
	}

	buf := make([]byte, 128)
	n, err := stream.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Client: Got size: %d, msg: '%s'", n, string(buf[:n]))
}

func do2() {
	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			//RootCAs:            pool,
			InsecureSkipVerify: true,
			//KeyLogWriter:       keyLog,
		},
		QuicConfig: &quic.Config{},
	}
	defer roundTripper.Close()

	cli := &http.Client{
		Transport: roundTripper,
	}

	resp, err := cli.Get("https://127.0.0.1:8586/hello")
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(string(data))
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	do2()
}
