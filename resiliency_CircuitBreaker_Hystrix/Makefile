proto:
	cd shared-proto \
	&& protoc --go_out=plugins=grpc:../client-svc/pb product.proto \
	&& protoc --go_out=plugins=grpc:../product-svc/pb product.proto

start:
	cd product-svc && go run main.go