package domain

import "context"

type TodoRepository interface {
	GetTodos(ctx context.Context, query *GetTodosQuery) ([]*Todo, error)
}
