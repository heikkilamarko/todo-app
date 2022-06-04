package internal

import (
	"context"
	"database/sql"
	_ "embed"
)

var (
	//go:embed sql/create_todo.sql
	createTodoSQL string
	//go:embed sql/complete_todo.sql
	completeTodoSQL string
)

type PostgresRepository struct {
	db *sql.DB
}

func (r *PostgresRepository) CreateTodo(ctx context.Context, todo *Todo) error {
	return r.db.QueryRowContext(ctx, createTodoSQL,
		todo.Name,
		todo.Description,
		todo.CreatedAt,
		todo.UpdatedAt).
		Scan(&todo.ID)
}

func (r *PostgresRepository) CompleteTodo(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx, completeTodoSQL, id)
	return err
}
