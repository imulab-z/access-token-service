FROM golang:1.11
RUN mkdir -p /go/src/github.com/imulab-z/access-token-service
ADD . /go/src/github.com/imulab-z/access-token-service
WORKDIR /go/src/github.com/imulab-z/access-token-service
RUN CGO_ENABLED=0 GOOS=linux go build -o access-token-service .

FROM alpine:3.9.2
RUN apk --no-cache add ca-certificates
WORKDIR /bin
COPY --from=0 /go/src/github.com/imulab-z/access-token-service/access-token-service .

CMD ["access-token-service"]