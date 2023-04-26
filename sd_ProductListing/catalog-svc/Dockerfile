# Dockerfile
FROM golang:1.17-alpine as builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o main

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .

CMD ["./main"]