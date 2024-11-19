package service

import (
	"todo-app/models"
	"todo-app/repository"
)

type TodosServiceInterface interface {
	FindAll() ([]*models.Todo, error)
	Create(todo models.Todo) (uint, error)
	Update(todo models.Todo) (models.Todo, error)
	Delete(id uint) error
}

type TodoService struct {
	repo repository.TodoRepositoryInterface
}

func NewTodoService(repo repository.TodoRepositoryInterface) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) FindAll() ([]*models.Todo, error) {
	return s.repo.FindAll()
}

func (s *TodoService) Create(todo models.Todo) (uint, error) {
	return s.repo.Create(todo)
}

func (s *TodoService) Update(todo models.Todo) (models.Todo, error) {
	return s.repo.Update(todo)
}

func (s *TodoService) Delete(id uint) error {
	return s.repo.Delete(id)
}
