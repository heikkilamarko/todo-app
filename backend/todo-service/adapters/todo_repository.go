package adapters

import (
	"context"
	"database/sql"
	_ "embed"
	"time"
	"todo-service/domain"
)

var (
	//go:embed sql/create_todo.sql
	createTodoSQL string
	//go:embed sql/complete_todo.sql
	completeTodoSQL string
)

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db}
}

func (r *TodoRepository) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	now := time.Now()

	todo.CreatedAt = now
	todo.UpdatedAt = now

	err := r.db.QueryRowContext(ctx, createTodoSQL,
		todo.Name,
		todo.Description,
		todo.CreatedAt,
		todo.UpdatedAt).
		Scan(&todo.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *TodoRepository) CompleteTodo(ctx context.Context, id int) error {
	if _, err := r.db.ExecContext(ctx, completeTodoSQL, id); err != nil {
		return err
	}
	return nil
}
