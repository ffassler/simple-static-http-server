FROM alpine:3.9
RUN apk add --no-cache libc6-compat
COPY simple-http-server .
EXPOSE 8080/tcp
CMD ["./simple-http-server"]
