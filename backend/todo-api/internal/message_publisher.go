package internal

import "context"

type MessagePublisher interface {
	TodoCreate(ctx context.Context, todo *Todo) error
	TodoComplete(ctx context.Context, id int) error
}
