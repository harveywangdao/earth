package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/proxy"
	"h12.io/socks"
)

func do1() {
	dialer, err := proxy.SOCKS5("tcp", "localhost:8080", nil, proxy.Direct)
	if err != nil {
		log.Fatal(err)
		return
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
	}

	resp, err := httpClient.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	log.Printf("%s", body)
}

func do2() {
	dialSocksProxy := socks.Dial("socks5://127.0.0.1:8080?timeout=5s")
	tr := &http.Transport{Dial: dialSocksProxy}
	httpClient := &http.Client{Transport: tr}
	resp, err := httpClient.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("%s", body)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
	do2()
}
