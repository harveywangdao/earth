package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
)

func do1() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_ = rdb.FlushDB(context.Background()).Err()

	limiter := redis_rate.NewLimiter(rdb)
	res, err := limiter.Allow(context.Background(), "project:123", redis_rate.PerSecond(10))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Allowed, res.Remaining)

	res, err = limiter.Allow(context.Background(), "project:123", redis_rate.PerSecond(10))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Allowed, res.Remaining)

	time.Sleep(2 * time.Second)

	res, err = limiter.Allow(context.Background(), "project:123", redis_rate.PerSecond(10))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Allowed, res.Remaining)

	res, err = limiter.Allow(context.Background(), "project:1234", redis_rate.PerSecond(10))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Allowed, res.Remaining)

	res, err = limiter.Allow(context.Background(), "project:1234", redis_rate.PerSecond(5))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Allowed, res.Remaining)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
