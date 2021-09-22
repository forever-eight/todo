package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/forever-eight/todo.git/internal/app/ds"
)

type Authorisation interface {
	CreateUser(user ds.User) (int, error)
	GetUser(username, password string) (ds.User, error)
}

type TodoList interface {
	Create(id int, list ds.TodoList) (int, error)
	GetAll(id int) ([]ds.TodoList, error)
	GetById(userId, listId int) (ds.TodoList, error)
	Delete(userId, id int) error
	Update(userId, id int, input ds.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, input ds.TodoItem) (int, error)
}

type Repository struct {
	Authorisation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
		TodoList:      NewToDoListPostgres(db),
	}
}
