package main

import (
	"net"
	"fmt"
	"log"
	"google.golang.org/grpc"

	"todo_proto/pb/proto"
	"todo_grpc/service"
	"google.golang.org/grpc/reflection"
)


func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterTodoServiceServer(grpcServer, &service.Server{})
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
