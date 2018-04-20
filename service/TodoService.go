package service

import (
	"golang.org/x/net/context"
	"todo_proto/pb/proto"
	db2 "todo_grpc/db"
	"strconv"
)

type Server struct {
	dao *db2.TodoDao
}

func NewServer() *Server {
	return &Server{db2.NewTodoDao()}
}

func (s *Server) CreateTodo(ctx context.Context, createTodoRequest *proto.CreateTodoRequest) (*proto.CreateTodoResponse, error) {

	todo := createTodoRequest.Todo

	rowsAffected, err := s.dao.CreateTodo(todo)
	if (err != nil) {
		return nil, err
	} else {
		return &proto.CreateTodoResponse{Id: strconv.Itoa(int(rowsAffected))}, nil
	}

}

func (s *Server) GetTodo(ctx context.Context, request *proto.GetTodoRequest) (*proto.GetTodoResponse, error) {
	todo, err := s.dao.GetTodo(request.Id)
	if (err != nil) {
		return nil, err
	} else {
		return &proto.GetTodoResponse{Todo: &todo}, nil
	}
}
