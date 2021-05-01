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

//go:embed sql/get_todos.sql
var qetTodosSQL string

// GetTodos method
func (r *SQLRepository) GetTodos(ctx context.Context, query *GetTodosQuery) ([]*Todo, error) {
	rows, err := r.db.QueryContext(
		ctx,
		qetTodosSQL,
		query.Limit, query.Offset)

	if err != nil {
		r.logger.Err(err).Send()
		return nil, goutils.ErrInternalError
	}

	defer rows.Close()

	todos := []*Todo{}

	for rows.Next() {
		t := &Todo{}

		err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.Description,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			r.logger.Err(err).Send()
			return nil, goutils.ErrInternalError
		}

		todos = append(todos, t)
	}

	return todos, nil
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
