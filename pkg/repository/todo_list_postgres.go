package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	todo "github.com/ohpasha/rest-api"
	"github.com/sirupsen/logrus"
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
		if err := tx.Rollback(); err != nil {
			return 0, err
		}

		return 0, err
	}

	queryUserList := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2) returning id", usersListsTable)

	_, err = tx.Exec(queryUserList, userId, id)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return 0, err
		}

		return 0, err
	}

	return id, tx.Commit()
}

func (r *TodolistPostgres) GetAll(usedId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList
	query := fmt.Sprintf("SELECT tl.id,tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1", todoListTable, usersListsTable)
	err := r.db.Select(&lists, query, usedId)

	return lists, err
}

func (r *TodolistPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf(`SELECT tl.id,tl.title, tl.description FROM %s tl INNER JOIN %s ul
		on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2`, todoListTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

func (r *TodolistPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id=$1 AND ul.list_id=$2", todoListTable, usersListsTable)

	_, err := r.db.Exec(query, userId, listId)

	return err
}

func (r *TodolistPostgres) Update(userId, listId int, input todo.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id =$%d",
		todoListTable, setQuery, usersListsTable, argId, argId+1)

	args = append(args, listId, userId)

	logrus.Debugf("update query: %s", setQuery)
	logrus.Debugf("args: %s", args...)

	_, err := r.db.Exec(query, args...)

	return err
}
