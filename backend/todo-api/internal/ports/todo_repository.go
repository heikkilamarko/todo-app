package ports

import (
	"context"
	"todo-api/internal/domain"
)

type GetTodosQuery struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type TodoRepository interface {
	GetTodos(ctx context.Context, query *GetTodosQuery) ([]*domain.Todo, error)
}
