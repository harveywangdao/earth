package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func do1(cli *clientv3.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	putResp, err := cli.Put(ctx, "sample_key", "sample_value")
	cancel()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%+v\n", *putResp)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	getResp, err := cli.Get(ctx, "sample_key")
	cancel()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%+v\n", *getResp)

	for _, ev := range getResp.Kvs {
		log.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func do2(cli *clientv3.Client) {
	for i := 0; i < 10; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		_, err := cli.Put(ctx, fmt.Sprintf("sample_key_%d", i), "sample_value")
		cancel()
		if err != nil {
			log.Println(err)
			return
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	getResp, err := cli.Get(ctx, "sample_key", clientv3.WithPrefix())
	cancel()
	if err != nil {
		log.Println(err)
		return
	}

	for _, ev := range getResp.Kvs {
		log.Printf("%s : %s\n", ev.Key, ev.Value)
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
	log.Printf("%+v\n", *cli)

	//do1(cli)
	do2(cli)
}
