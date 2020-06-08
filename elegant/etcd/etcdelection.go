package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err = el.Campaign(ctx, "thomas")
	cancel()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("elected")

	go do2(cli)

	time.Sleep(5 * time.Second)

	getResp, err := el.Leader(context.Background())
	if err != nil {
		log.Println(err)
		return
	}
	for _, ev := range getResp.Kvs {
		log.Printf("%s : %s\n", ev.Key, ev.Value)
	}

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

	time.Sleep(20 * time.Second)
}

func do2(cli *clientv3.Client) {
	session, err := concurrency.NewSession(cli)
	if err != nil {
		log.Println(err)
		return
	}

	el := concurrency.NewElection(session, "darwin")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = el.Campaign(ctx, "thomas2")
	cancel()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("do2 elected")
	time.Sleep(10 * time.Second)
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
