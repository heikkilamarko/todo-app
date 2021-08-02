package adapters

import (
	"context"
	"database/sql"
	_ "embed"
	"time"
	"todo-service/internal/domain"
)

var (
	//go:embed sql/create_todo.sql
	createTodoSQL string
	//go:embed sql/complete_todo.sql
	completeTodoSQL string
)

type TodoPostgresRepository struct {
	db *sql.DB
}

func NewTodoPostgresRepository(db *sql.DB) *TodoPostgresRepository {
	return &TodoPostgresRepository{db}
}

func (r *TodoPostgresRepository) CreateTodo(ctx context.Context, todo *domain.Todo) error {
	now := time.Now()

	todo.CreatedAt = now
	todo.UpdatedAt = now

	return r.db.QueryRowContext(ctx, createTodoSQL,
		todo.Name,
		todo.Description,
		todo.CreatedAt,
		todo.UpdatedAt).
		Scan(&todo.ID)
}

func (r *TodoPostgresRepository) CompleteTodo(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, completeTodoSQL, id)
	return err
}
