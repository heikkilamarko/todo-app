package todos

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/heikkilamarko/goutils"
	"github.com/rs/zerolog"
)

type repository struct {
	db     *sql.DB
	logger *zerolog.Logger
}

func newRepository(db *sql.DB, logger *zerolog.Logger) *repository {
	return &repository{db, logger}
}

//go:embed sql/create_todo.sql
var createTodoSQL string

func (r *repository) createTodo(ctx context.Context, command *createTodoCommand) error {
	t := command.Todo

	n := time.Now()

	t.CreatedAt = n
	t.UpdatedAt = n

	err := r.db.QueryRowContext(ctx, createTodoSQL,
		t.Name,
		t.Description,
		t.CreatedAt,
		t.UpdatedAt).
		Scan(&t.ID)

	if err != nil {
		r.logger.Err(err).Send()
		return goutils.ErrInternalError
	}

	return nil
}
