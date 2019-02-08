FROM golang:1.11 as builder
COPY main.go .
RUN GOOS=linux go build -o simple-http-server

FROM alpine:3.9
RUN apk add --no-cache libc6-compat
COPY --from=0 /go/simple-http-server .
EXPOSE 8080/tcp
CMD ["./simple-http-server"]
