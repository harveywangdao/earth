package main

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
)

func main() {
	config := clientv3.Config{
		Endpoints: []string{"http://192.168.1.10:2764"},
	}
	c, err := clientv3.New(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(c)
}
