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
}

type TodoItem interface {
}

type Repository struct {
	Authorisation
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorisation: NewAuthPostgres(db),
	}
}
