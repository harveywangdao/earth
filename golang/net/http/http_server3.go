package main

import (
	"log"
	"net/http"
	"time"
)

func do1() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		log.Println("sleep start")
		time.Sleep(time.Second * 5)
		log.Println("sleep end")

		w.Write([]byte("hello"))
	})

	log.Println("http server listen:", "8564")
	log.Fatal(http.ListenAndServe(":8564", mux))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
