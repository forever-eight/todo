package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"

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

func (r *ToDoListPostgres) GetAll(userId int) ([]ds.TodoList, error) {
	var lists []ds.TodoList
	// tl* - все поля с tl
	query := fmt.Sprintf("SELECT tl.* FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1",
		todoListsTable, usersListsTable)
	err := r.db.Select(&lists, query, userId)
	return lists, err
}

func (r *ToDoListPostgres) GetById(userId, listId int) (ds.TodoList, error) {
	var list ds.TodoList
	query := fmt.Sprintf("SELECT tl.id, tl.title, tl.description FROM %s tl INNER JOIN %s ul on tl.id = ul.list_id WHERE ul.user_id = $1 AND ul.list_id = $2",
		todoListsTable, usersListsTable)
	err := r.db.Get(&list, query, userId, listId)
	fmt.Println(err)
	return list, err
}

func (r *ToDoListPostgres) Delete(userId, id int) error {
	query := fmt.Sprintf("DELETE FROM %s tl USING %s ul WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2", todoListsTable, usersListsTable)
	res, err := r.db.Exec(query, userId, id)

	if res == nil {
		fmt.Println("empty")
		return fmt.Errorf("empty db")
	}
	return err
}

func (r *ToDoListPostgres) Update(userId, id int, input ds.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1
	if input.Title != nil {
		// Добавляем значение арг сюда, то есть $1
		setValues = append(setValues, fmt.Sprintf("title = $%d", argID))
		// Сюда добавляем то, что хотим положить
		args = append(args, *input.Title)
		// Увеличиваем
		argID++
	}
	if input.Description != nil {
		// Добавляем значение арг сюда, то есть $1
		setValues = append(setValues, fmt.Sprintf("description = $%d", argID))
		// Сюда добавляем то, что хотим положить
		args = append(args, *input.Description)
		// Увеличиваем
		argID++
	}
	// Соединяем строки. Пример
	// title=$1
	// title=$1, description=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
		todoListsTable, setQuery, usersListsTable, argID, argID+1)
	args = append(args, id, userId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
