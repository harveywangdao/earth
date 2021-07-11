package main

import (
	"io/ioutil"
	"log"
	"net/http"
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

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
