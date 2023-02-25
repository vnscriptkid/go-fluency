package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	pb "server-svc/hello"

	"google.golang.org/grpc"
)

type helloServer struct{}

func (server *helloServer) HelloSomeone(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	switch req.Language {
	case "english":
		return &pb.HelloResponse{
			Greeting: fmt.Sprintf("Hello %s %s", req.Person.FirstName, req.Person.LastName),
		}, nil
	case "vietnamese":
		return &pb.HelloResponse{
			Greeting: fmt.Sprintf("Xin chao %s %s", req.Person.LastName, req.Person.FirstName),
		}, nil
	default:
		return nil, errors.New("unsupported language " + req.Language)
	}
}

func main() {
	port := "50051"

	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("[server] fail to listen %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterHelloServiceServer(s, &helloServer{})

	log.Println("[server] started at :" + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("[server] failed to serve %v", err)
	}
}
