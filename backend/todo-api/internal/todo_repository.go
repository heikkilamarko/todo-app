package internal

import "context"

type GetTodosQuery struct {
	Offset int
	Limit  int
}

type GetTodosResult struct {
	Todos []Todo
}

type TodoRepository interface {
	GetTodos(ctx context.Context, q *GetTodosQuery) (*GetTodosResult, error)
}
