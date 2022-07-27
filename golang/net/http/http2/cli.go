package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func do1() {
	client := &http.Client{
		Transport: &http.Transport{
			ForceAttemptHTTP2: true,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	req, err := http.NewRequest(http.MethodGet, "https://localhost:8564", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

	/*	req.Proto = "HTTP/2.0"
		req.ProtoMajor = 2
		req.ProtoMinor = 0
	*/
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	log.Println(resp.Status)
	log.Println(resp.Proto)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
