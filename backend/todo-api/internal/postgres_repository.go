package internal

import (
	"context"
	"database/sql"
	_ "embed"
)

//go:embed sql/get_todos.sql
var getTodosSQL string

type PostgresRepository struct {
	DB *sql.DB
}

func (r *PostgresRepository) GetTodos(ctx context.Context, q *GetTodosQuery) ([]*Todo, error) {
	rows, err := r.DB.QueryContext(
		ctx,
		getTodosSQL,
		q.Limit, q.Offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := []*Todo{}

	for rows.Next() {
		d := &Todo{}

		err := rows.Scan(
			&d.ID,
			&d.Name,
			&d.Description,
			&d.CreatedAt,
			&d.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		data = append(data, d)
	}

	return data, nil
}
