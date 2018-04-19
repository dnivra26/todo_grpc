package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"todo_proto/pb/proto"
)

func TestNewTodoDao_GetTodo(t *testing.T) {

	dao := NewTodoDao()
	db := dao.db
	defer db.Close()
	todo := proto.Todo{Id: "20", Title: "second todo", Status: true}
	dao.CreateTodo(&todo)

	todo2 := dao.GetTodo(todo.Id)

	assert.Equal(t, todo, todo2)

}

func TestNewTodoDao_CreateTodo(t *testing.T) {
	dao := NewTodoDao()
	db := dao.db
	defer db.Close()
	todo := proto.Todo{Id: "22", Title: "second todo", Status: true}
	rowsAffected := dao.CreateTodo(&todo)
	assert.Equal(t, 1, rowsAffected)
}
