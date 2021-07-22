package service

import "github.com/forever-eight/todo.git/internal/app/repository"

type Authorisation interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorisation
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
