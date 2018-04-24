package db

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/dnivra26/todo_proto/pb/proto"
)

func TestNewTodoDao_GetTodo(t *testing.T) {

	dao := NewTodoDao()
	db := dao.db
	defer db.Close()
	todo := proto.Todo{Id: "262314660", Title: "second todo", Status: true}
	dao.CreateTodo(&todo)

	todo2, _ := dao.GetTodo("262314660")

	assert.Equal(t, todo, todo2)

}

func TestNewTodoDao_CreateTodo(t *testing.T) {
	dao := NewTodoDao()
	db := dao.db
	defer db.Close()
	todo := proto.Todo{Id: "232", Title: "second todo", Status: true}
	rowsAffected, _ := dao.CreateTodo(&todo)
	assert.Equal(t, 1, rowsAffected)
}
