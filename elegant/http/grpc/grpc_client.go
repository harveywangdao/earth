package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/harveywangdao/earth/elegant/http/grpc/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//ctx = metadata.NewIncomingContext(ctx, metadata.Pairs("Incomingkey01", "Incomingvalue01"))
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("Outgoingkey01", "Outgoingvalue01"))

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMsg())
}
