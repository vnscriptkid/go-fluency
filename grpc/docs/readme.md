## snippets
- `cd common-lib`
- `protoc --go_out=plugins=grpc:../client-svc/hello hello.proto`
- `protoc --go_out=plugins=grpc:../server-svc/hello hello.proto`
- `cd server-svc`
- `mockery --all --output=./mocks`