package internal

import "context"

type Repository interface {
	GetPermissions(ctx context.Context, roles []string) ([]string, error)
	GetTodos(ctx context.Context, q *GetTodosQuery) ([]*Todo, error)
}
