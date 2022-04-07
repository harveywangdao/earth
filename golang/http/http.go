package main

import (
	"fmt"
	"log"
	"net/http"
)

// curl http://localhost:8080
func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	log.Println("RequestURI:", r.RequestURI)
	fmt.Fprintf(w, "Hi, This is an example of http service in golang\n")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	http.HandleFunc("/", handler)
	log.Println("listening http server")
	http.ListenAndServe(":8080", nil)
}
