package service

import (
	"golang.org/x/net/context"
	"todo_proto/pb/proto"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Server struct {
	db *gorm.DB
}

//type Todo struct {
//	gorm.Model
//	Id string
//	Title string
//	Status bool
//}

func (s *Server) CreateTodo(ctx context.Context, createTodoRequest *proto.CreateTodoRequest) (*proto.CreateTodoResponse, error) {
	todo := createTodoRequest.Todo
	create := s.db.Create(todo)
	return &proto.CreateTodoResponse{Id: strconv.Itoa(int(create.RowsAffected))}, nil
}

func (s *Server) GetTodo(ctx context.Context, request *proto.GetTodoRequest) (*proto.GetTodoResponse, error) {
	var todo proto.Todo
	s.db.First(&todo, request.Id)
	return &proto.GetTodoResponse{Todo: &todo}, nil
}
