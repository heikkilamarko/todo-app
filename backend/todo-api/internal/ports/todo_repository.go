package ports

import (
	"context"
	"todo-api/internal/domain"
)

type GetTodosQuery struct {
	Offset int
	Limit  int
}

type TodoRepository interface {
	GetTodos(ctx context.Context, query *GetTodosQuery) ([]*domain.Todo, error)
}
