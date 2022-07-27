package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func do1() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header)
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("http server recv msg: %s", data)
		w.Write([]byte("hello, I am http server"))
		//w.WriteHeader(http.StatusOK)
	})

	log.Println("http server listen:", "8564")
	log.Fatal(http.ListenAndServe(":8564", mux))
}

func do2() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header)
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("http server recv msg: %s", data)

		//w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Connection", "close")
		w.Write([]byte("hello, I am http server"))
	})

	log.Println("http server listen:", "8564")
	log.Fatal(http.ListenAndServe(":8564", mux))
}

func do3() {
	mux := http.NewServeMux()
	mux.HandleFunc("/chunk", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header)

		buf := make([]byte, 128)
		for {
			n, err := r.Body.Read(buf)
			if err != nil {
				if err == io.EOF {
					if n > 0 {
						log.Printf("http server recv msg: %s", buf[:n])
					}
					break
				}
				log.Fatal(err)
			}
			log.Printf("http server recv msg: %s", buf[:n])
		}

		//w.Header().Set("Transfer-Encoding", "chunked")
		w.Write([]byte("apple1"))
		w.(http.Flusher).Flush()
		time.Sleep(time.Second)

		w.Write([]byte("apple2"))
		w.(http.Flusher).Flush()
		time.Sleep(time.Second)

		w.Write([]byte("apple3"))
		w.(http.Flusher).Flush()
		time.Sleep(time.Second)

		w.Write([]byte("apple4"))
		w.(http.Flusher).Flush()
		time.Sleep(time.Second)

		w.Write([]byte("apple5"))
		w.(http.Flusher).Flush()
	})

	log.Println("http server listen:", "8564")
	log.Fatal(http.ListenAndServe(":8564", mux))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do3()
}
