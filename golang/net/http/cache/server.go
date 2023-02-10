package main

import (
	"log"
	"net/http"
)

func do1() {
	mux := http.NewServeMux()
	mux.HandleFunc("/dog", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Header)
		data := make([]byte, 1024*1024*50)
		for i := 0; i < len(data); i++ {
			data[i] = byte(i%26) + 'a'
		}
		w.Write(data)
		w.WriteHeader(http.StatusOK)
	})

	log.Println("http server listen:", "8564")
	log.Fatal(http.ListenAndServe(":8564", mux))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
