package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	pb "random-svc/random"

	"google.golang.org/grpc"
)

type randomServer struct{}

func (server *randomServer) Random(ctx context.Context, req *pb.RandomRequest) (*pb.RandomResponse, error) {
	randInt := rand.Intn(int(req.To)-int(req.From)+1) + int(req.From)

	return &pb.RandomResponse{
		Result: int32(randInt),
	}, nil
}

func main() {
	port := "50052"

	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("[server] fail to listen %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterRandomServiceServer(s, &randomServer{})

	log.Println("[server] started at :" + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("[server] failed to serve %v", err)
	}
}
