FROM golang:alpine AS builder
RUN apk add --no-cache --update git
ENV GOPATH=/go

ADD . $GOPATH/src/todo_grpc
WORKDIR $GOPATH/src/todo_grpc

RUN go get

RUN go build -v -o /go/bin/todo_grpc server.go

FROM alpine:latest

COPY --from=builder /go/bin/todo_grpc /usr/local/bin/todo_grpc

CMD todo_grpc