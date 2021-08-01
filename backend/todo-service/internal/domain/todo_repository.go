package domain

import "context"

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *Todo) error
	CompleteTodo(ctx context.Context, id int) error
}
