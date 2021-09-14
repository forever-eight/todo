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

func (s *ToDoListService) GetById(userId, listId int) (ds.TodoList, error) {
	return s.r.GetById(userId, listId)
}

func (s *ToDoListService) Delete(userId, id int) error {
	return s.r.Delete(userId, id)
}

func (s *ToDoListService) Update(userId, id int, input ds.UpdateListInput) error {
	err := input.Validate()
	if err != nil {
		return err
	}
	return s.r.Update(userId, id, input)
}
