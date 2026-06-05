package service

import "go-supabase-todos/model"

type TodoService interface {
    GetAll() ([]model.Todo, error)
    GetByID(id int) (*model.Todo, error)
    Create(req model.CreateTodoRequest) (*model.Todo, error)
    Update(id int, req model.UpdateTodoRequest) (*model.Todo, error)
    Delete(id int) error
}