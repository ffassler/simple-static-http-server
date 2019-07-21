FROM golang:1.12.7 as builder
COPY main.go .
COPY main_test.go .
RUN GOOS=linux go build -o simple-http-server
RUN GOOS=linux go get github.com/stretchr/testify/assert
RUN GOOS=linux go test

FROM alpine:3.9
RUN apk add --no-cache libc6-compat
COPY --from=0 /go/simple-http-server .
EXPOSE 8080/tcp
CMD ["./simple-http-server"]
