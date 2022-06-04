package internal

import "context"

type Repository interface {
	CreateTodo(ctx context.Context, todo *Todo) error
	CompleteTodo(ctx context.Context, id int) error
}
