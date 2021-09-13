package service

import (
	"github.com/forever-eight/todo.git/internal/app/ds"
	"github.com/forever-eight/todo.git/internal/app/repository"
)

type Authorisation interface {
	CreateUser(user ds.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(id int, list ds.TodoList) (int, error)
	GetAll(id int) ([]ds.TodoList, error)
	GetById(userId, listId int) (ds.TodoList, error)
	Delete(userId, id int) error
}

type TodoItem interface {
}

type Service struct {
	Authorisation
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorisation: NewAuthService(repos.Authorisation),
		TodoList:      NewToDoListService(repos.TodoList),
	}

}
