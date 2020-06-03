package main

import (
	"fmt"
	"net/http"
)

/*
openssl genrsa -out server.key 2048
openssl req -new -x509 -key server.key -out server.crt -days 365
*/

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,
		"Hi, This is an example of https service in golang!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServeTLS(":8081", "server.crt", "server.key", nil)
}
