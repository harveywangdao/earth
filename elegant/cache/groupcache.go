package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"net/rpc"
	"strconv"
	"time"

	"github.com/golang/groupcache"
)

type Store struct {
	Key   string
	Value string
}

type Bird struct {
	cacheGroup *groupcache.Group
}

func (b *Bird) Get(req *Store, resp *Store) error {
	var data []byte
	err := b.cacheGroup.Get(context.Background(), req.Key, groupcache.AllocatingByteSliceSink(&data))
	resp.Value = string(data)
	return err
}

func server(port int) {
	peers := groupcache.NewHTTPPool("http://localhost:" + strconv.Itoa(port))
	peers.Set("http://localhost:8001", "http://localhost:8002", "http://localhost:8003")

	stringcache := groupcache.NewGroup("SlowDBCache", 64<<20, groupcache.GetterFunc(
		func(ctx context.Context, key string, dest groupcache.Sink) error {
			result := "nasklxnakm"
			time.Sleep(time.Second * 3)
			log.Println("set key from db to cache:", key)
			dest.SetBytes([]byte(result))
			return nil
		}))

	go func() {
		bird := &Bird{cacheGroup: stringcache}
		rpc.Register(bird)
		rpc.HandleHTTP()
		http.ListenAndServe(":"+strconv.Itoa(port+1000), nil)
	}()

	http.ListenAndServe(":"+strconv.Itoa(port), http.HandlerFunc(peers.ServeHTTP))
}

func client(key string) {
	client, err := rpc.DialHTTP("tcp", "localhost:9001")
	if err != nil {
		log.Fatal(err)
		return
	}

	req := &Store{
		Key: key,
	}
	resp := &Store{}
	err = client.Call("Bird.Get", req, resp)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(resp.Value)
}

/*
./cache -key key02
./cache -server -port 8001
./cache -server -port 8002
./cache -server -port 8003

*/
func do1() {
	var port = flag.Int("port", 8001, "groupcache port")
	var bserver = flag.Bool("server", false, "server or client")
	var key = flag.String("key", "key01", "key")
	flag.Parse()

	if *bserver {
		server(*port)
	} else {
		client(*key)
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
