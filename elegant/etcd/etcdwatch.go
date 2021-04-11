package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
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

	go do2(cli)

	for i := 0; i < 10; i++ {
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		putResp, err = cli.Put(ctx, fmt.Sprintf("sample_key/%d", i), fmt.Sprintf("sample_value%d", i))
		cancel()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%+v\n", *putResp)
		time.Sleep(time.Second)

		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		delResp, err := cli.Delete(ctx, fmt.Sprintf("sample_key/%d", i))
		cancel()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%+v\n", *delResp)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	getResp, err := cli.Get(ctx, "sample_key11")
	cancel()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%+v\n", *getResp)

	for _, ev := range getResp.Kvs {
		log.Printf("%s : %s\n", ev.Key, ev.Value)
	}

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	delResp, err := cli.Delete(ctx, "sample_key")
	cancel()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%+v\n", *delResp)

	time.Sleep(time.Minute)
}

func do2(cli *clientv3.Client) {
	rch := cli.Watch(context.Background(), "sample_key", clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			log.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
	log.Println("watch exit")
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
