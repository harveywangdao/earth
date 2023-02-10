package main

import (
	"log"
	"net/http"
)

func do1() {
	resp, err := http.DefaultClient.Get("http://192.168.126.134:8564/dog")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println(resp.Header)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
