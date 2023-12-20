package main

import (
	"context"
	"net"

	"google.golang.org/grpc"

	pb "my-svc/pkg/api"
)

type server struct {
	pb.UnimplementedYourServiceServer
}

func (s *server) Echo(ctx context.Context, in *pb.StringMessage) (*pb.StringMessage, error) {
	return in, nil
}

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterYourServiceServer(grpcServer, &server{})
	lis, _ := net.Listen("tcp", ":50051")
	grpcServer.Serve(lis)
}
