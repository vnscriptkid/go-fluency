package main

import (
	pb "client-svc/hello"
	"client-svc/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
)

func Test_DoSomething(t *testing.T) {
	mockHelloServiceClient := &mocks.HelloServiceClient{}
	injectedRes := &pb.HelloResponse{
		Greeting: "Example",
	}

	mockHelloServiceClient.On("HelloSomeone", mock.Anything, &pb.HelloRequest{
		Person: &pb.Person{
			FirstName: "Example fName",
			LastName:  "Example lName",
		},
		Language: "Example lang",
	}).Return(injectedRes, nil).Once()

	server := Server{
		clients: &serviceClients{
			helloServiceClient: mockHelloServiceClient,
		},
	}

	res, err := server.SayHello("Example fName", "Example lName", "Example lang")

	if err != nil {
		t.Error("Expect err is nil")
	}

	expectedRes := "Response: greeting:\"Example\" "

	if res != expectedRes {
		t.Errorf("Expected `%v` but got `%v`", expectedRes, res)
	}
}
