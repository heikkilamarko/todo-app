package internal

import "context"

type MessagePublisher interface {
	TodoCreateOk(ctx context.Context, todo *Todo) error
	TodoCreateError(ctx context.Context, message string) error
	TodoCompleteOk(ctx context.Context, id int) error
	TodoCompleteError(ctx context.Context, message string) error
}
