package main

import (
	"log"
	"net/http"

	"github.com/didip/tollbooth"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello, World!"))
}

func do1() {
	http.Handle("/", tollbooth.LimitFuncHandler(tollbooth.NewLimiter(1, nil), HelloHandler))
	http.ListenAndServe(":12345", nil)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
