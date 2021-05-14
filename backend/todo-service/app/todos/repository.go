package todos

import (
	"context"
	"database/sql"
	"time"

	"github.com/heikkilamarko/goutils"
	"github.com/rs/zerolog"
)

type repository struct {
	db     *sql.DB
	logger *zerolog.Logger
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
		r.logger.Error().Err(err).Send()
		return goutils.ErrInternalError
	}

	return nil
}

func (r *repository) completeTodo(ctx context.Context, command *completeTodoCommand) error {
	_, err := r.db.ExecContext(ctx, sqlCompleteTodo, command.ID)

	if err != nil {
		r.logger.Err(err).Send()
		return goutils.ErrInternalError
	}

	return nil
}
