package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	todo "github.com/ohpasha/rest-api"
)

type AuthPostges struct {
	db *sqlx.DB
}

func NewAuthPostges(db *sqlx.DB) *AuthPostges {
	return &AuthPostges{db: db}
}

func (r *AuthPostges) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) returning id", userTable)
	row := r.db.QueryRow(query, user.Name, user.UserName, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostges) GetUser(username, password string) (todo.User, error) {
	var user todo.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
