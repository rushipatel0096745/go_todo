package service

import (
    "go-supabase-todos/model"
    "go-supabase-todos/repository"
)

type todoService struct {
    repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) TodoService {
    return &todoService{repo: repo}
}

func (s *todoService) GetAll() ([]model.Todo, error) {
    return s.repo.FindAll()
}

func (s *todoService) GetByID(id int) (*model.Todo, error) {
    return s.repo.FindByID(id)
}

func (s *todoService) Create(req model.CreateTodoRequest) (*model.Todo, error) {
    return s.repo.Create(req)
}

func (s *todoService) Update(id int, req model.UpdateTodoRequest) (*model.Todo, error) {
    return s.repo.Update(id, req)
}

func (s *todoService) Delete(id int) error {
    return s.repo.Delete(id)
}