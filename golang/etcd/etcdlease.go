package main

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
)

func do1(cli *clientv3.Client) {
	resp, err := cli.Grant(context.Background(), 10)
	if err != nil {
		log.Println(err)
		return
	}

	keepAliveCh, err := cli.KeepAlive(context.Background(), resp.ID)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		leaseKeepAliveResp, ok := <-keepAliveCh
		if !ok {
			log.Println("keepAliveCh ch closed")
			return
		}
		log.Println(leaseKeepAliveResp.ID, leaseKeepAliveResp.TTL)

		timeToLiveResp, err := cli.TimeToLive(context.Background(), resp.ID)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(timeToLiveResp.ID, timeToLiveResp.TTL, timeToLiveResp.GrantedTTL, timeToLiveResp.Keys)
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
