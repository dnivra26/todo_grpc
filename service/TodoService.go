package service

import (
	"golang.org/x/net/context"
	"todo_proto/pb/proto"
	db2 "todo_grpc/db"
	"strconv"
	"google.golang.org/grpc"
	"log"
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
	todo.Title, _ = getHello()
	if (err != nil) {
		return nil, err
	} else {
		return &proto.GetTodoResponse{Todo: &todo}, nil
	}
}

func getHello() (string, error) {
	conn, e := grpc.Dial(":7778", grpc.WithInsecure())
	if e != nil {
	log.Fatalf("did not connect: %s", e)
	}
	defer conn.Close()
	client := proto.NewPingClient(conn)
	response, err := client.SayHello(context.Background(), &proto.PingMessage{Greeting: "Hello"})
	return response.Greeting, err
}
