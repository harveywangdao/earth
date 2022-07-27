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

	log.Println("https 2.0 server listen:", "8564")
	log.Fatal(http.ListenAndServeTLS(":8564", "../certificate/server.crt", "../certificate/server.key", mux))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
