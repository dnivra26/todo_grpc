package db

import (
	"github.com/dnivra26/todo_proto/pb/proto"
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

type TodoDao struct {
	db *gorm.DB
}

func getDB() (*gorm.DB) {
	dbDSN := fmt.Sprintf("user=%s password=%s host=%s DB.name=gtd dbname=gtd port=5432 sslmode=disable", "gtd", "password", os.Getenv("DB_HOST"))
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
func (todoDao *TodoDao) CreateTodo(todo *proto.Todo) (int, error) {
	result := todoDao.db.Create(&todo)
	err := result.Error
	return int(result.RowsAffected), err
}

func (todoDao *TodoDao) GetTodo(id string) (proto.Todo, error) {
	var todo proto.Todo
	result := todoDao.db.First(&todo, id)
	err := result.Error
	return todo, err
}
