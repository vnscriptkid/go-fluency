package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/grpc"

	pb "client-svc/pb/server-svc"
)

const (
	address = "localhost:50051"
)

func main() {
	// Load image data
	imageData, err := ioutil.ReadFile("photo1.jpg")

	if err != nil {
		log.Fatalf("failed to read image file %v", err)
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("failed to connect %v", err)
	}

	defer conn.Close()

	c := pb.NewImageServiceClient(conn)

	stream, err := c.UploadImage(context.Background())

	err = stream.Send(&pb.Image{
		Name: "photo1.jpg",
		Data: imageData,
	})

	if err != nil {
		log.Fatalf("err sending image %v", err)
	}

	_, err = stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("err closing stream %v", err)
	}

	fmt.Println("Image uploaded successfully")
}
