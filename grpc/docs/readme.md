## snippets
- `cd common-lib`
- `protoc --go_out=plugins=grpc:../client-svc/hello hello.proto`
- `protoc --go_out=plugins=grpc:../server-svc/hello hello.proto`
- `cd server-svc`
- `mockery --all --output=./mocks`

## ginkgo
- Install v1: `go install github.com/onsi/ginkgo/ginkgo@latest`
- New executable created `cd $GOPATH/bin`
- Check `which ginkgo`
- Version `ginkgo version`
- VS Code support `vscode-ginkgo`
- Gen boilerplate `ginkgo bootstrap`

## gomock
- Use mockgen go gen mocks: 
    - `cd client-svc`
    - `mockgen --destination ./gomocks/hello.go --source=./hello/hello.pb.go --build_flags=--mod=mod`

## Deps
- See what required a lib: `go mod why -m github.com/golang/mock`