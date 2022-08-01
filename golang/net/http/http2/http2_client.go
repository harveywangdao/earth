package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"sync"
	"time"
)

func httpReq(cli *http.Client, i int) {
	req, err := http.NewRequest(http.MethodGet, "https://127.0.0.1:8564/api", nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("req start", i)
	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println("req end", i, resp.Proto)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("http2.0 client recv msg: %s, code: %v", data, resp.StatusCode)
}

func do1() {
	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:   true,
			MaxIdleConns:        100,
			IdleConnTimeout:     90 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
			MaxIdleConnsPerHost: 2,
			MaxConnsPerHost:     2,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			httpReq(cli, index)
		}(i)
	}

	wg.Wait()
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
