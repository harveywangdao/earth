package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/harveywangdao/earth/golang/logger/logger"
	pb "github.com/harveywangdao/earth/golang/net/grpc/hello"
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

	ctx = logger.CreateLogContext(ctx)

	logger.Info("start client")
	logger.Infof("%s", "start client")

	logger.With(ctx).Info("start client")
	logger.With(ctx).Infof("%s", "start client")

	ctx = logger.CreateOutgoingContext(ctx)
	ctx = metadata.AppendToOutgoingContext(ctx, "x-out-going", "1287319823791", "x-qws-asdas", "dasdasdasd")

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMsg())
}
