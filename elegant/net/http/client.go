package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func do1() {
	buf := bytes.NewBufferString("hello, I am http client")
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:8564/api", buf)
	if err != nil {
		log.Fatal(err)
	}
	cli := &http.Client{}

	resp, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	log.Println(resp.Header)

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("http client recv msg: %s", data)

	if resp.StatusCode != http.StatusOK {
		log.Fatal("http code:", resp.Status)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
