package main

import (
	"log"
	"net/http"
)

func do1() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Proto, r.ProtoMajor, r.ProtoMinor)
		log.Println(r.URL)
	})

	log.Println("http2 server listen:", "8564")
	log.Fatal(http.ListenAndServe(":8564", mux))
}

func main() {
	do1()
}
