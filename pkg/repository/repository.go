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
}

type TodoItem interface {
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
	}
}
