package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/forever-eight/todo.git/internal/app/ds"
)

type ToDoListPostgres struct {
	db *sqlx.DB
}

func NewToDoListPostgres(db *sqlx.DB) *ToDoListPostgres {
	return &ToDoListPostgres{db: db}
}

func (r *ToDoListPostgres) Create(userId int, list ds.TodoList) (int, error) {
	// Начало транзакции
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListsTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	err = row.Scan(&id)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return 0, err
		}
		return 0, err
	}
	return id, tx.Commit()
}
