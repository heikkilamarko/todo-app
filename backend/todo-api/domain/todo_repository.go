package domain

import "context"

type TodoRepository interface {
	GetTodos(ctx context.Context, query *GetTodosQuery) ([]*Todo, error)
	CreateTodo(ctx context.Context, todo *Todo) error
	CompleteTodo(ctx context.Context, id int) error
}
