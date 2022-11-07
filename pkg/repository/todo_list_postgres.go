package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	todo "github.com/ohpasha/rest-api"
)

type TodolistPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgers(db *sqlx.DB) *TodolistPostgres {
	return &TodolistPostgres{
		db: db,
	}
}

func (r *TodolistPostgres) Create(userId int, todoList todo.TodoList) (int, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var id int

	queryList := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) returning id", todoListTable)
	row := tx.QueryRow(queryList, todoList.Title, todoList.Description)

	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	queryUserList := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2) returning id", usersListsTable)

	_, err = tx.Exec(queryUserList, userId, id)

	if err != nil {
		tx.Rollback()

		return 0, err
	}

	return id, tx.Commit()
}
