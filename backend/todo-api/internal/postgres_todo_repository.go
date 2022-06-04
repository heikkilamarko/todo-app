package internal

import (
	"context"
	"database/sql"
	_ "embed"
)

//go:embed sql/get_todos.sql
var getTodosSQL string

type PostgresTodoRepository struct {
	db *sql.DB
}

func NewPostgresTodoRepository(db *sql.DB) *PostgresTodoRepository {
	return &PostgresTodoRepository{db}
}

func (r *PostgresTodoRepository) GetTodos(ctx context.Context, q *GetTodosQuery) (*GetTodosResult, error) {
	rows, err := r.db.QueryContext(
		ctx,
		getTodosSQL,
		q.Limit, q.Offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []Todo{}

	for rows.Next() {
		var d Todo

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

		todos = append(todos, d)
	}

	return &GetTodosResult{todos}, nil
}
