package ports

import (
	"context"
	"todo-service/internal/domain"
)

type TodoRepository interface {
	CreateTodo(ctx context.Context, todo *domain.Todo) error
	CompleteTodo(ctx context.Context, id int) error
}
