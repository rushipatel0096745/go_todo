package repository

import (
	"database/sql"
	"fmt"
	"go-supabase-todos/model"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (r *PostgresRepo) FindAll() ([]model.Todo, error) {
	rows, err := r.db.Query(
		"SELECT id, title, completed FROM todos ORDER BY id",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (r *PostgresRepo) FindByID(id int) (*model.Todo, error) {
	var todo model.Todo
	err := r.db.QueryRow(
		"SELECT id, title, completed FROM todos WHERE id = $1", id,
	).Scan(&todo.ID, &todo.Title, &todo.Completed)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("todo not found")
	}
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *PostgresRepo) Create(req model.CreateTodoRequest) (*model.Todo, error) {
	var todo model.Todo
	err := r.db.QueryRow(
		"INSERT INTO todos (title) VALUES ($1) RETURNING id, title, completed",
		req.Title,
	).Scan(&todo.ID, &todo.Title, &todo.Completed)

	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *PostgresRepo) Update(id int, req model.UpdateTodoRequest) (*model.Todo, error) {
	var todo model.Todo
	err := r.db.QueryRow(
		"UPDATE todos SET title = $1, completed = $2 WHERE id = $3 RETURNING id, title, completed",
		req.Title, req.Completed, id,
	).Scan(&todo.ID, &todo.Title, &todo.Completed)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("todo not found")
	}
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *PostgresRepo) Delete(id int) error {
	result, err := r.db.Exec(
		"DELETE FROM todos WHERE id = $1", id,
	)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("todo not found")
	}
	return nil
}
