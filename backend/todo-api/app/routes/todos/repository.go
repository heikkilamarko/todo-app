package todos

import (
	"context"
	"database/sql"
	_ "embed"

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

//go:embed sql/get_todos.sql
var qetTodosSQL string

func (r *repository) getTodos(ctx context.Context, query *getTodosQuery) ([]*todo, error) {
	rows, err := r.db.QueryContext(
		ctx,
		qetTodosSQL,
		query.Limit, query.Offset)

	if err != nil {
		r.logger.Err(err).Send()
		return nil, goutils.ErrInternalError
	}

	defer rows.Close()

	todos := []*todo{}

	for rows.Next() {
		t := &todo{}

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
