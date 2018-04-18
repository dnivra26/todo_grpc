package main

import (
	"net"
	"fmt"
	"log"
	"google.golang.org/grpc"

	"todo_proto/pb/proto"
	"todo_grpc/service"
	"google.golang.org/grpc/reflection"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	dbDSN := fmt.Sprintf("user=%s password=%s DB.name=gtd dbname=gtd port=5432 sslmode=disable", "gtd", "password")
	db, err := gorm.Open("postgres", dbDSN)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&proto.Todo{})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterTodoServiceServer(grpcServer, &service.Server{db})
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
