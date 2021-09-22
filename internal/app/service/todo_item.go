package service

import (
	"github.com/forever-eight/todo.git/internal/app/ds"
	"github.com/forever-eight/todo.git/internal/app/repository"
)

type ToDoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewToDoItemService(repo repository.TodoItem) *ToDoItemService {
	return &ToDoItemService{repo: repo}
}
func (s *ToDoItemService) Create(userId, listId int, input ds.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		// list does not exist
		return 0, err
	}

	return s.repo.Create(userId, listId, input)
}
