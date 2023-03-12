package main

import (
	mocks "client-svc/gomocks"
	pb "client-svc/hello"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_GoMocks(t *testing.T) {
	var errExpected error = nil
	resultExpected := "Response: greeting:\"example only\" "

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockHelloServiceClient(ctrl)

	mockClient.EXPECT().HelloSomeone(gomock.Any(), gomock.Any()).Return(&pb.HelloResponse{
		Greeting: "example only",
	}, nil).AnyTimes()

	server := Server{
		clients: &serviceClients{
			helloServiceClient: mockClient,
		},
	}

	result, err := server.SayHello("x", "y", "z")

	if err != errExpected {
		t.Errorf("Expected %v but got %v", errExpected, err)
	}

	if result != resultExpected {
		t.Errorf("Expected %v but got %v", resultExpected, result)
	}
}

func Test_GoMocks_2(t *testing.T) {
	var errExpected error = nil
	resultExpected := "Response: greeting:\"example only\" "

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockHelloServiceClient(ctrl)

	mockClient.EXPECT().HelloSomeone(gomock.Any(), gomock.Any()).
		DoAndReturn(func(_ context.Context, _ *pb.HelloRequest) (*pb.HelloResponse, error) {
			/*
				In case of: github.com/golang/mock v1.6.0
				Should be:
				.HelloSomeone(gomock.Any(), gomock.Any(), gomock.Any()).
				DoAndReturn(func(_ context.Context, _ *pb.HelloRequest, opts ...grpc.CallOption) (*pb.HelloResponse, error) {
			*/
			return &pb.HelloResponse{
				Greeting: "example only",
			}, nil
		})

	server := Server{
		clients: &serviceClients{
			helloServiceClient: mockClient,
		},
	}

	result, err := server.SayHello("x", "y", "z")

	if err != errExpected {
		t.Errorf("Expected %v but got %v", errExpected, err)
	}

	if result != resultExpected {
		t.Errorf("Expected %v but got %v", resultExpected, result)
	}
}
