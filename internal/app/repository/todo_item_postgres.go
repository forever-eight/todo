package repository

import "github.com/jmoiron/sqlx"

type ToDoItemPostgres struct {
	DB *sqlx.DB
}

func newToDoItemPostgres(DB *sqlx.DB) *ToDoItemPostgres {
	return &ToDoItemPostgres{DB: DB}
}
