package main

import (
	"context"
	"log"
	"net"
	pb "product-svc/pb"
	"time"

	"google.golang.org/grpc"
)

type productServer struct{}

func (s *productServer) GetProduct(context.Context, *pb.ProductRequest) (*pb.ProductResponse, error) {
	// Simulate a long running process by random sleep from 500ms to 2s
	randomDuration := time.Duration(500+time.Now().UnixNano()%1300) * time.Millisecond

	log.Printf("[product-svc] processing took %v", randomDuration)

	time.Sleep(randomDuration)

	return &pb.ProductResponse{
		Product: &pb.Product{
			Id:          "1",
			Name:        "Product 1",
			Price:       1000,
			Description: "",
			Category:    "",
			Image:       "",
		},
	}, nil

}

func main() {
	port := "50051"

	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("[product-svc] fail to listen %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterProductServiceServer(s, &productServer{})

	log.Println("[product-svc] started at :" + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("[product-svc] failed to serve %v", err)
	}
}
