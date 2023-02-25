package main

import (
	pb "client-svc/hello"
	"context"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("[client] failed to connect %v", err)
	}

	defer conn.Close()

	c := pb.NewHelloServiceClient(conn)

	req := &pb.HelloRequest{
		Person: &pb.Person{
			FirstName: "Thanh",
			LastName:  "Nguyen",
		},
		Language: "english",
	}

	res, err := c.HelloSomeone(context.Background(), req)

	if err != nil {
		log.Fatalf("failed to hello %v", err)
	}

	log.Printf("hello response %v", res)

}
