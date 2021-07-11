package main

import (
	"context"
	"log"
	"time"

	pb "github.com/harveywangdao/earth/elegant/grpc/hello"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func unaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("grpc client method: %s cost: %s", method, time.Since(start))
	return err
}

func do1() {
	conn, err := grpc.DialContext(context.Background(), address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithUnaryInterceptor(unaryClientInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	cli := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	resp, err := cli.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("say hello fail: %v", err)
		return
	}
	log.Println(resp)
}

// 不公用tcp链接
func do2() {
	num := 10
	ch := make(chan int, num)
	for {
		ch <- 1
		go func() {
			do1()
			<-ch
		}()
	}
}

// 公用一个tcp链接
func do3() {
	conn, err := grpc.DialContext(context.Background(), address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithUnaryInterceptor(unaryClientInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	cli := pb.NewGreeterClient(conn)

	num := 10
	ch := make(chan int, num)
	for {
		ch <- 1
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
			defer cancel()

			resp, err := cli.SayHello(ctx, &pb.HelloRequest{Name: "world"})
			if err != nil {
				log.Fatalf("say hello fail: %v", err)
			}
			log.Println(resp)
			<-ch
		}()
	}
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	do3()
}
