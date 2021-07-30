package domain

import "context"

type TodoMessagePublisher interface {
	CreateTodo(ctx context.Context, todo *Todo) error
	CompleteTodo(ctx context.Context, id int) error
}
