package main

import (
	"context"
	"log"
	"net"

	pb "github.com/harveywangdao/earth/elegant/http/grpc/hello"
	"github.com/harveywangdao/earth/elegant/logger/logger"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/metadata"
)

const (
	port = ":50051"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	//md, ok := metadata.FromIncomingContext(ctx)
	//log.Println(md, ok)
	/*md, ok = metadata.FromOutgoingContext(ctx)
	log.Println(md, ok)*/

	ctx = logger.CreateLogContext(ctx)

	logger.Info("Received:", in.GetName())
	logger.Infof("Received: %v", in.GetName())

	logger.With(ctx).Info("Received:", in.GetName())
	logger.With(ctx).Infof("Received: %v", in.GetName())

	return &pb.HelloReply{Msg: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	logger.Info("start service")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
