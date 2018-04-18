package service

import (
	"golang.org/x/net/context"
	"todo_proto/pb/proto"
)

type Server struct {
}

func (s *Server) CreateTodo(ctx context.Context, createTodoRequest *proto.CreateTodoRequest) (*proto.CreateTodoResponse, error) {
	return &proto.CreateTodoResponse{Id: "1"}, nil
}

func (s *Server) GetTodo(ctx context.Context, request *proto.GetTodoRequest) (*proto.GetTodoResponse, error) {
	return &proto.GetTodoResponse{Todo: &proto.Todo{"", "", true}}, nil
}

