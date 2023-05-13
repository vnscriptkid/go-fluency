package main

import (
	"context"
	"log"

	"client-svc/pb"

	"github.com/afex/hystrix-go/hystrix"
	"google.golang.org/grpc"
)

func main() {
	// Initialize a gRPC connection
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Initialize your gRPC client
	productClient := pb.NewProductServiceClient(conn)

	// Configure the circuit breaker
	hystrix.ConfigureCommand("GetProduct", hystrix.CommandConfig{
		// Timeout in milliseconds, if GetProduct takes longer than this, request will timeout
		Timeout: 1000,
		// maximum of 50 requests can be processed concurrently
		MaxConcurrentRequests: 50,
		// Consider last 20 requests for error rate calculation
		// If 50% or more of the requests fail, the circuit will open
		ErrorPercentThreshold: 50,
		// If less than 20 requests are processed, the circuit will not open even if all requests fail
		// We need large enough sample to draw a reliable conclusion
		RequestVolumeThreshold: 20,
		// Give 5 seconds for the circuit to recover before trying again
		// After 5 seconds, switch to HALF_OPEN state, allow a single request to pass through for a trial
		// If trial succeeds, the circuit will close again
		// If trial fails, the circuit will open again for another 5 seconds
		SleepWindow: 5000,
	})

	output := make(chan *pb.ProductResponse, 1)
	errors := hystrix.Go("GetProduct", func() error {
		// Do your RPC call
		product, err := productClient.GetProduct(context.Background(), &pb.ProductRequest{ProductId: "1"})
		if err != nil {
			return err
		}

		// Send the response to the output channel
		output <- product
		return nil
	}, nil)

	select {
	case out := <-output:
		log.Printf("Response: %v", out)
	case err := <-errors:
		log.Printf("Error: %v", err)
	}

}

// hystrix.ConfigureCommand("service1_call", hystrix.CommandConfig{
// 	Timeout:                500,
// 	MaxConcurrentRequests:  50,
// 	RequestVolumeThreshold: 3,
// 	SleepWindow:            2000,
// 	ErrorPercentThreshold:  25,
// })

// hystrix.ConfigureCommand("service2_call", hystrix.CommandConfig{
// 	Timeout:                1000,
// 	MaxConcurrentRequests:  100,
// 	RequestVolumeThreshold: 5,
// 	SleepWindow:            5000,
// 	ErrorPercentThreshold:  50,
// })
