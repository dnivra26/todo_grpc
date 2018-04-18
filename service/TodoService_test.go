package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"todo_proto/pb/proto"
)

func TestServer_GetTodo(t *testing.T) {
	server := Server{}

	response, error := server.GetTodo(context.Background(), &proto.GetTodoRequest{Id: "1"})

	assert.Equal(t, "", response.Todo.Id)
	assert.Equal(t, "", response.Todo.Title)
	assert.Equal(t, true, response.Todo.Status)

	assert.Nil(t, error)

}

func TestServer_CreateTodo(t *testing.T) {
	server := Server{}
	todo := proto.Todo{Id: "1", Title: "", Status: true}
	response, error := server.CreateTodo(context.Background(), &proto.CreateTodoRequest{Todo: &todo})

	assert.Equal(t, "1", response.Id)
	assert.Nil(t, error)
}
