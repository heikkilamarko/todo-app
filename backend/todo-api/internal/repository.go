package internal

import "context"

type Repository interface {
	GetTodos(ctx context.Context, q *GetTodosQuery) ([]*Todo, error)
}
