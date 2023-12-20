#!/bin/bash
protoc -I ../api -I ../googleapis \
  --go_out=plugins=grpc:../pkg/api \
  ../api/api.proto

protoc -I ../api -I ../googleapis \
  --grpc-gateway_out=logtostderr=true:../pkg/api \
  ../api/api.proto
