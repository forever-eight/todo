package service

import (
	"github.com/forever-eight/todo.git/internal/app/ds"
	"github.com/forever-eight/todo.git/internal/app/repository"
)

type ToDoListService struct {
	r repository.TodoList
}

func NewToDoListService(r repository.TodoList) *ToDoListService {
	return &ToDoListService{r: r}
}

func (s *ToDoListService) Create(id int, list ds.TodoList) (int, error) {
	return s.r.Create(id, list)
}

func (s *ToDoListService) GetAll(id int) ([]ds.TodoList, error) {
	return s.r.GetAll(id)
}

func (s ToDoListService) GetById(userId, listId int) (ds.TodoList, error) {
	return s.r.GetById(userId, listId)
}
