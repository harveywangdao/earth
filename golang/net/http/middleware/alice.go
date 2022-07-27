package main

import (
	"log"
	"net/http"
	"time"

	"github.com/justinas/alice"
	"github.com/justinas/nosurf"
	"github.com/throttled/throttled"
	"github.com/throttled/throttled/store/memstore"
)

func timeoutHandler(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 1*time.Second, "timed out")
}

func myApp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func main() {
	store, err := memstore.New(65536)
	if err != nil {
		log.Fatal(err)
	}

	quota := throttled.RateQuota{
		MaxRate:  throttled.PerMin(1),
		MaxBurst: 1,
	}
	rateLimiter, err := throttled.NewGCRARateLimiter(store, quota)
	if err != nil {
		log.Fatal(err)
	}

	httpRateLimiter := throttled.HTTPRateLimiter{
		RateLimiter: rateLimiter,
		VaryBy:      &throttled.VaryBy{Path: true},
	}

	myHandler := http.HandlerFunc(myApp)

	chain := alice.New(httpRateLimiter.RateLimit, timeoutHandler, nosurf.NewPure).Then(myHandler)
	http.ListenAndServe(":8000", chain)
}
