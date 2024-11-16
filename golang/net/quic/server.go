package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"

	quic "github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
)

func do1() {
	generateTLSConfig := func() *tls.Config {
		key, err := rsa.GenerateKey(rand.Reader, 1024)
		if err != nil {
			panic(err)
		}
		template := x509.Certificate{SerialNumber: big.NewInt(1)}
		certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
		if err != nil {
			panic(err)
		}
		keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
		certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

		tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
		if err != nil {
			panic(err)
		}
		return &tls.Config{
			Certificates: []tls.Certificate{tlsCert},
			NextProtos:   []string{"quic-echo-example"},
		}
	}

	listener, err := quic.ListenAddr(":8586", generateTLSConfig(), nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		sess, err := listener.Accept(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		stream, err := sess.AcceptStream(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		buf := make([]byte, 128)
		n, err := stream.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Server: Got size: %d, msg: '%s'", n, string(buf[:n]))

		msg := string(buf[:n]) + " haha"
		_, err = stream.Write([]byte(msg))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func do2() {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Println(err)
		return
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		log.Println(err)
		return
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	ioutil.WriteFile("priv.key", keyPEM, os.ModePerm)
	ioutil.WriteFile("cert.pem", certPEM, os.ModePerm)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("http3 request")
		w.Write([]byte("http3 response"))
	})

	log.Fatal(http3.ListenAndServe(":8586", "cert.pem", "priv.key", mux))
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	do2()
}
