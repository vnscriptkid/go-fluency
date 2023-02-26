package main

import (
	"context"
	"errors"
	"reflect"
	pb "server-svc/hello"
	"testing"
)

func Test_Errors(t *testing.T) {
	err1 := errors.New("my error")
	err2 := errors.New("my error")
	err3 := errors.New("her error")

	if !reflect.DeepEqual(err1, err2) {
		t.Error("Expect the same errors")
	}

	if reflect.DeepEqual(err1, err3) {
		t.Error("Expect different errors")
	}
}

func Test_HelloSomeone(t *testing.T) {
	server := helloServer{}

	testcases := []struct {
		name        string
		helloReq    *pb.HelloRequest
		expectedRes *pb.HelloResponse
		expectedErr error
	}{
		{
			name: "TC1",
			helloReq: &pb.HelloRequest{
				Language: "vietnamese",
				Person: &pb.Person{
					FirstName: "thanh",
					LastName:  "nguyen",
				},
			},
			expectedRes: &pb.HelloResponse{
				Greeting: "Xin chao nguyen thanh",
			},
			expectedErr: nil,
		},
		{
			name: "TC2",
			helloReq: &pb.HelloRequest{
				Language: "english",
				Person: &pb.Person{
					FirstName: "Thanh",
					LastName:  "Nguyen",
				},
			},
			expectedRes: &pb.HelloResponse{
				Greeting: "Hello Thanh Nguyen",
			},
			expectedErr: nil,
		},
		{
			name: "TC3",
			helloReq: &pb.HelloRequest{
				Language: "xxx",
				Person: &pb.Person{
					FirstName: "Thanh",
					LastName:  "Nguyen",
				},
			},
			expectedRes: nil,
			expectedErr: errors.New("unsupported language xxx"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := server.HelloSomeone(context.Background(), tc.helloReq)

			if !reflect.DeepEqual(err, tc.expectedErr) {
				t.Errorf("Expected err is `%v` but got `%v`", tc.expectedErr, err)
			}

			if res == nil && tc.expectedRes == nil {
				return
			}

			if !reflect.DeepEqual(*res, *tc.expectedRes) {
				t.Errorf("Expected res is %v but got %v", *tc.expectedRes, *res)
			}
		})
	}
}
