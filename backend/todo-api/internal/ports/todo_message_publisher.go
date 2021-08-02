package ports

import (
	"context"
	"todo-api/internal/domain"
)

type TodoMessagePublisher interface {
	TodoCreate(ctx context.Context, todo *domain.Todo) error
	TodoComplete(ctx context.Context, id int) error
}
