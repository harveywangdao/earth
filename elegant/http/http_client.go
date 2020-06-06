package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	client := &http.Client{}

	resp, err := client.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println("http request not 200")
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("body:", string(data))
}
