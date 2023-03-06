package main

import (
	"io"
	"log"
	"net"
	"os"

	pb "server-svc/pb/server-svc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) UploadImage(stream pb.ImageService_UploadImageServer) error {
	imageData := []byte{}
	imageName := ""

	for {
		// data from the stream, which represents the client's image data being sent in chunks
		req, err := stream.Recv()

		if err == io.EOF {
			// Finished reading the image
			break
		}

		if err != nil {
			return status.Errorf(codes.Unknown, "err reading image %v", err)
		}

		imageData = append(imageData, req.GetData()...)
		imageName = req.GetName()
	}

	file, err := os.Create(imageName)

	if err != nil {
		return status.Errorf(codes.Internal, "err creating file %v", err)
	}

	defer file.Close()

	_, err = file.Write(imageData)

	if err != nil {
		return status.Errorf(codes.Internal, "error writing to file %v", err)
	}

	return stream.SendAndClose(&pb.ImageResponse{
		Message: "Done",
	})
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterImageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}
