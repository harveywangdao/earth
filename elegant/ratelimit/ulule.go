package main

import (
	"log"
	"net/http"

	"github.com/ulule/limiter"
	"github.com/ulule/limiter/drivers/middleware/stdlib"
	"github.com/ulule/limiter/drivers/store/memory"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write([]byte(`{"message": "ok"}`))
	if err != nil {
		log.Fatal(err)
	}
}

func do1() {
	rate, err := limiter.NewRateFromFormatted("1-S")
	if err != nil {
		log.Fatal(err)
		return
	}

	store := memory.NewStore()
	middleware := stdlib.NewMiddleware(limiter.New(store, rate, limiter.WithTrustForwardHeader(true)))

	http.Handle("/", middleware.Handler(http.HandlerFunc(index)))
	log.Println("Server is running on port 7777...")
	log.Fatal(http.ListenAndServe(":7777", nil))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
