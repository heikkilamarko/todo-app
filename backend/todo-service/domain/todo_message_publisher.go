package domain

import "context"

type TodoMessagePublisher interface {
	TodoCreatedOk(ctx context.Context, todo *Todo) error
	TodoCreatedError(ctx context.Context, message string) error
	TodoCompletedOk(ctx context.Context, id int) error
	TodoCompletedError(ctx context.Context, message string) error
}
