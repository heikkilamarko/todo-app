package todos

import (
	"context"
	"database/sql"
	"time"
)

type repository struct {
	db *sql.DB
}

func (r *repository) createTodo(ctx context.Context, command *createTodoCommand) error {
	t := command.Todo

	n := time.Now()

	t.CreatedAt = n
	t.UpdatedAt = n

	err := r.db.QueryRowContext(ctx, sqlCreateTodo,
		t.Name,
		t.Description,
		t.CreatedAt,
		t.UpdatedAt).
		Scan(&t.ID)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) completeTodo(ctx context.Context, command *completeTodoCommand) error {

	if _, err := r.db.ExecContext(ctx, sqlCompleteTodo, command.ID); err != nil {
		return err
	}

	return nil
}
