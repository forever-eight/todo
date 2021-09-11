package service

import (
	"github.com/forever-eight/todo.git/internal/app/ds"
	"github.com/forever-eight/todo.git/internal/app/repository"
)

type ToDoListService struct {
	r *repository.Repository
}

func NewToDoListService(r *repository.Repository) *ToDoListService {
	return &ToDoListService{r: r}
}

func (s ToDoListService) Create(id int, list ds.TodoList) (int, error) {
	return s.r.Create(id, list)
}