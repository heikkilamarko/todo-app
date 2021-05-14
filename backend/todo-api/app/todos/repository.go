package todos

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

func (r *repository) getTodos(ctx context.Context, query *getTodosQuery) ([]*todo, error) {
	rows, err := r.db.QueryContext(
		ctx,
		sqlGetTodos,
		query.Limit, query.Offset)

	if err != nil {
		return nil, err
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
			return nil, err
		}

		todos = append(todos, t)
	}

	return todos, nil
}
