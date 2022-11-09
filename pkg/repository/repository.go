package repository

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/ohpasha/rest-api"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(name, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, todoList todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, id int) error
	Update(userId, id int, input todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostges(db),
		TodoList:      NewTodoListPostgers(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
