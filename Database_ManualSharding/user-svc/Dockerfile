FROM golang:1.17 as builder

WORKDIR /app
COPY . .

# RUN go mod init example.com/sharding && go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

FROM scratch
COPY --from=builder /app/main /main

ENTRYPOINT ["/main"]
