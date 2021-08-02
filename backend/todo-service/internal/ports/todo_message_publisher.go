package ports

import (
	"context"
	"todo-service/internal/domain"
)

type TodoMessagePublisher interface {
	TodoCreateOk(ctx context.Context, todo *domain.Todo) error
	TodoCreateError(ctx context.Context, message string) error
	TodoCompleteOk(ctx context.Context, id int) error
	TodoCompleteError(ctx context.Context, message string) error
}
