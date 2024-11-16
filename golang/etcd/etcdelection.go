package main

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
	"log"
	"time"
)

func do1(cli *clientv3.Client) {
	session, err := concurrency.NewSession(cli)
	if err != nil {
		log.Println(err)
		return
	}

	el := concurrency.NewElection(session, "darwin")
	go do3(el)
	time.Sleep(time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err = el.Campaign(ctx, "thomas")
	cancel()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("do1 elected")

	go do2(cli)

	getResp, err := el.Leader(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	for _, ev := range getResp.Kvs {
		log.Printf("%s : %s\n", ev.Key, ev.Value)
	}

	log.Println("do1 sleeping 10s...")
	time.Sleep(10 * time.Second)

	err = el.Resign(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Resigned")
	time.Sleep(5 * time.Second)

	getResp, err = el.Leader(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	for _, ev := range getResp.Kvs {
		log.Printf("%s : %s\n", ev.Key, ev.Value)
	}

	time.Sleep(2000 * time.Second)
}

func do2(cli *clientv3.Client) {
	session, err := concurrency.NewSession(cli)
	if err != nil {
		log.Println(err)
		return
	}

	el := concurrency.NewElection(session, "darwin")
	ctx, cancel := context.WithTimeout(context.Background(), 12*time.Second)
	log.Println("do2 waiting 12s...")
	err = el.Campaign(ctx, "thomas2")
	cancel()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("do2 elected")
	time.Sleep(1000 * time.Second)
}

func do3(el *concurrency.Election) {
	o := el.Observe(context.Background())

	for {
		select {
		case resp, ok := <-o:
			if !ok {
				log.Println("Observe fail")
				return
			}

			log.Println("current master:", resp)
		}
	}
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
