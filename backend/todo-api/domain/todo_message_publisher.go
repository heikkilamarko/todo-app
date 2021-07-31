package domain

import "context"

type TodoMessagePublisher interface {
	TodoCreate(ctx context.Context, todo *Todo) error
	TodoComplete(ctx context.Context, id int) error
}
