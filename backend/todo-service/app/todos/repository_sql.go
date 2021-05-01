package todos

import (
	"context"
	"database/sql"
	_ "embed"
	"time"

	"github.com/heikkilamarko/goutils"
	"github.com/rs/zerolog"
)

// SQLRepository struct
type SQLRepository struct {
	db     *sql.DB
	logger *zerolog.Logger
}

// NewSQLRepository func
func NewSQLRepository(db *sql.DB, l *zerolog.Logger) *SQLRepository {
	return &SQLRepository{db, l}
}

//go:embed sql/create_todo.sql
var createTodoSQL string

// CreateTodo method
func (r *SQLRepository) CreateTodo(ctx context.Context, command *CreateTodoCommand) error {
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
