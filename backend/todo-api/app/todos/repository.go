package todos

import (
	"context"
	"database/sql"
	"todo-api/app/utils"

	"github.com/rs/zerolog"
)

type repository struct {
	db     *sql.DB
	logger *zerolog.Logger
}

func (r *repository) getTodos(ctx context.Context, query *getTodosQuery) ([]*todo, error) {
	rows, err := r.db.QueryContext(
		ctx,
		sqlGetTodos,
		query.Limit, query.Offset)

	if err != nil {
		r.logger.Error().Err(err).Send()
		return nil, utils.ErrInternalError
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
			r.logger.Error().Err(err).Send()
			return nil, utils.ErrInternalError
		}

		todos = append(todos, t)
	}

	return todos, nil
}
