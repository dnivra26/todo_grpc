package db

import (
	"todo_proto/pb/proto"
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type TodoDao struct {
	db *gorm.DB
}

func getDB() (*gorm.DB) {
	dbDSN := fmt.Sprintf("user=%s password=%s DB.name=gtd dbname=gtd port=5432 sslmode=disable", "gtd", "password")
	db, err := gorm.Open("postgres", dbDSN)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	return db
}

func NewTodoDao() *TodoDao {
	return &TodoDao{db: getDB()}
}
func (todoDao *TodoDao) CreateTodo(todo *proto.Todo) (int) {
	result := todoDao.db.Create(&todo)
	return int(result.RowsAffected)
}

func (todoDao *TodoDao) GetTodo(id string) (proto.Todo) {
	var todo proto.Todo
	todoDao.db.First(&todo, id)
	return todo
}
