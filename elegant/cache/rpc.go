package main

import (
	"log"
	"net/http"
	"net/rpc"
	"time"
)

type Adder struct{}
type AddReq struct {
	N, M int
}
type AddResp struct {
	Sum int
}

func (a *Adder) Add(req *AddReq, resp *AddResp) error {
	log.Println(req.N, req.M)
	resp.Sum = req.N + req.M

	return nil
}

func server() {
	s := rpc.NewServer()
	s.Register(&Adder{})
	s.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
	http.ListenAndServe(":8464", s)
}

func client() {
	c, err := rpc.DialHTTP("tcp", "127.0.0.1:8464")
	if err != nil {
		log.Fatal(err)
		return
	}

	req := &AddReq{
		N: 12,
		M: 23,
	}
	resp := &AddResp{}
	err = c.Call("Adder.Add", req, resp)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(resp.Sum)
}

func do1() {
	go server()

	time.Sleep(1 * time.Second)
	client()
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do1()
}
