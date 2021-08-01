package adapters

import (
	"context"
	"database/sql"
	_ "embed"
	"todo-api/internal/domain"
)

//go:embed sql/get_todos.sql
var getTodosSQL string

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{db}
}

func (r *TodoRepository) GetTodos(ctx context.Context, query *domain.GetTodosQuery) ([]*domain.Todo, error) {
	rows, err := r.db.QueryContext(
		ctx,
		getTodosSQL,
		query.Limit, query.Offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []*domain.Todo{}

	for rows.Next() {
		t := &domain.Todo{}

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
