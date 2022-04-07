package main

import (
	"context"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

func do1() {
	rdb1 := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:7000", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	rdb2 := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:7001", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	ctx := context.Background()
	key := "key" + strconv.Itoa(rand.Intn(10000))
	value := "value" + strconv.Itoa(rand.Intn(10000))
	if err := rdb1.Set(ctx, key, value, 0).Err(); err != nil {
		panic(err)
	}
	log.Println("set:", key, value)

	val, err := rdb2.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	log.Println("get:", key, val)
}

func do2() {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"127.0.0.1:9000", "127.0.0.1:9001", "127.0.0.1:9002"},
	})
	ctx := context.Background()
	log.Println(rdb.Ping(ctx))

	key := "key" + strconv.Itoa(rand.Intn(10000))
	value := "value" + strconv.Itoa(rand.Intn(10000))
	if err := rdb.Set(ctx, key, value, 0).Err(); err != nil {
		panic(err)
	}
	log.Println("set:", key, value)

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	log.Println("get:", key, val)
}

func do3() {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"127.0.0.1:7000", "127.0.0.1:7001", "127.0.0.1:7002", "127.0.0.1:8000", "127.0.0.1:8001", "127.0.0.1:8002"},
	})
	ctx := context.Background()
	log.Println(rdb.Ping(ctx))

	key := "key" + strconv.Itoa(rand.Intn(10000))
	value := "value" + strconv.Itoa(rand.Intn(10000))
	if err := rdb.Set(ctx, key, value, 0).Err(); err != nil {
		panic(err)
	}
	log.Println("set:", key, value)

	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	log.Println("get:", key, val)
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	rand.Seed(time.Now().Unix())
	do3()
}
