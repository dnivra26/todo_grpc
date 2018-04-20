package service

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"todo_proto/pb/proto"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func getDB() (*gorm.DB) {
	dbDSN := fmt.Sprintf("user=%s password=%s DB.name=gtd dbname=gtd port=5432 sslmode=disable", "gtd", "password")
	db, err := gorm.Open("postgres", dbDSN)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	return db
}
func TestServer_GetTodo(t *testing.T) {
	db := getDB()
	defer db.Close()

	server := NewServer()

	todo := proto.Todo{Id: "112341", Title: "second todo", Status: true}
	server.CreateTodo(context.Background(), &proto.CreateTodoRequest{Todo: &todo})

	response, error := server.GetTodo(context.Background(), &proto.GetTodoRequest{Id: todo.Id})

	assert.Equal(t, &todo, response.Todo)

	assert.Nil(t, error)

}

func TestServer_CreateTodo(t *testing.T) {
	db := getDB()
	defer db.Close()

	server := NewServer()
	todo := proto.Todo{Id: "176500", Title: "", Status: true}
	response, error := server.CreateTodo(context.Background(), &proto.CreateTodoRequest{Todo: &todo})

	assert.Equal(t, "1", response.Id)
	assert.Nil(t, error)
}
