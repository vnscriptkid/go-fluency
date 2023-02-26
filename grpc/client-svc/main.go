package main

import (
	pb "client-svc/hello"
	random "client-svc/random"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

type serviceClients struct {
	helloServiceClient  pb.HelloServiceClient
	randomServiceClient random.RandomServiceClient
}

func NewServiceClients() (*serviceClients, error) {
	helloConn, err := grpc.Dial(":50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("[client] failed to connect %v", err)
	}

	randomConn, err := grpc.Dial(":50052", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("[client] failed to connect %v", err)
	}

	helloClient := pb.NewHelloServiceClient(helloConn)
	randomClient := random.NewRandomServiceClient(randomConn)

	return &serviceClients{
		helloServiceClient:  helloClient,
		randomServiceClient: randomClient,
	}, nil
}

type Server struct {
	clients *serviceClients
}

func (s *Server) SayHello(fName string, lName string, lang string) (string, error) {
	req := &pb.HelloRequest{
		Person: &pb.Person{
			FirstName: fName,
			LastName:  lName,
		},
		Language: lang,
	}

	res, err := s.clients.helloServiceClient.HelloSomeone(context.Background(), req)

	if err != nil {
		return "", fmt.Errorf("failed to hello %v", err)
	}

	return fmt.Sprintf("Response: %v", res), nil
}

func main() {
	sClients, err := NewServiceClients()

	if err != nil {
		log.Fatal("oops")
	}

	server := Server{
		clients: sClients,
	}

	res, err := server.SayHello("Thanh", "Nguyen", "english")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(res)
}
