package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
openssl genrsa -out server.key 2048
openssl req -new -x509 -key server.key -out server.crt -days 365
openssl rsa -in server.key -out server.key.public

cd golang/src/earth/elegant/http
curl --insecure https://localhost:8081
curl --cacert certificate/server.crt  https://localhost:8081

openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -subj "/CN=abc.com" -days 5000 -out ca.crt
openssl genrsa -out server.key 2048
openssl req -new -key server.key -subj "/CN=localhost" -out server.csr
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 5000

openssl x509 -in 4033637_www.wangpear.top.pem -noout -text
openssl x509 -in 4033637_www.wangpear.top.pem -noout -subject
*/

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Method:", r.Method)
	log.Println("RequestURI:", r.RequestURI)
	fmt.Fprintf(w, "Hi, This is an example of https service in golang\n")
}

func server01() {
	http.HandleFunc("/", handler)
	log.Println("listening https server")
	http.ListenAndServeTLS(":8081", "certificate/server.crt", "certificate/server.key", nil)
}

func server02() {
	http.HandleFunc("/", handler)
	log.Println("listening https server")
	http.ListenAndServeTLS(":8081", "ca/server.crt", "ca/server.key", nil)
}

type myhandler struct {
}

func (h *myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, This is an example of http service in golang\n")
}

func server03() {
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile("ca/ca.crt")
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr:    ":8081",
		Handler: &myhandler{},
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	log.Println("listening https server")
	if err := s.ListenAndServeTLS("ca/server.crt", "ca/server.key"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	server02()
}
