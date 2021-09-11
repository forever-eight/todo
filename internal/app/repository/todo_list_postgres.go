package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/forever-eight/todo.git/internal/app/ds"
)

type ToDoListPostgres struct {
	db *sqlx.DB
}

func NewToDoListPostgres(db *sqlx.DB) *ToDoListPostgres {
	return &ToDoListPostgres{db: db}
}

func (r *ToDoListPostgres) Create(id int, list ds.TodoList) (int, error) {

}
