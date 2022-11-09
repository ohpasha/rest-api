package todo

import (
	"errors"
)

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (input UpdateListInput) Validate() error {

	if input.Description == nil && input.Title == nil {
		return errors.New("empty titile and description. Nothing to update")
	}

	return nil
}

func (input UpdateItemInput) Validate() error {

	if input.Description == nil && input.Title == nil && input.Done == nil {
		return errors.New("empty titile and description. Nothing to update")
	}

	return nil
}
