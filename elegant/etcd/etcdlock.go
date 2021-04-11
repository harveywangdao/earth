package main

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/clientv3/concurrency"
	"log"
	"time"
)

func do1(cli *clientv3.Client) {
	session, err := concurrency.NewSession(cli)
	if err != nil {
		log.Println(err)
		return
	}

	mu := concurrency.NewMutex(session, "darwin")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err = mu.Lock(ctx)
	cancel()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("locked")

	go do2(cli)
	time.Sleep(10 * time.Second)

	err = mu.Unlock(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("unlocked")
}

func do2(cli *clientv3.Client) {
	session, err := concurrency.NewSession(cli)
	if err != nil {
		log.Println(err)
		return
	}

	mu := concurrency.NewMutex(session, "darwin")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err = mu.Lock(ctx)
	cancel()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("do2 locked")
	time.Sleep(5 * time.Second)

	err = mu.Unlock(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("do2 unlocked")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	config := clientv3.Config{
		Endpoints: []string{"http://127.0.0.1:2379"},
	}
	cli, err := clientv3.New(config)
	if err != nil {
		log.Println(err)
		return
	}
	defer cli.Close()

	do1(cli)
}
