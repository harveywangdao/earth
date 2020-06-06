package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
)

func client1() {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Get("https://localhost:8081")
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

func client2() {
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile("ca/ca.crt")
	if err != nil {
		log.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{RootCAs: pool},
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("https://localhost:8081")
	if err != nil {
		log.Println("Get error:", err)
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

/*
openssl genrsa -out client.key 2048
openssl req -new -key client.key -subj "/CN=abcd.cn" -out client.csr
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 5000
openssl x509 -text -in client.crt -noout
*/
func client3() {
	pool := x509.NewCertPool()
	caCrt, err := ioutil.ReadFile("ca/ca.crt")
	if err != nil {
		log.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	cliCrt, err := tls.LoadX509KeyPair("ca/client.crt", "ca/client.key")
	if err != nil {
		log.Println(err)
		return
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs:      pool,
			Certificates: []tls.Certificate{cliCrt},
		},
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("https://localhost:8081")
	if err != nil {
		log.Println("Get error:", err)
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

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	client2()
}
