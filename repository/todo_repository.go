package repository

import "go-supabase-todos/model"

type TodoRepository interface {
    FindAll() ([]model.Todo, error)
    FindByID(id int) (*model.Todo, error)
    Create(req model.CreateTodoRequest) (*model.Todo, error)
    Update(id int, req model.UpdateTodoRequest) (*model.Todo, error)
    Delete(id int) error
}