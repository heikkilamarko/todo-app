package command

import (
	"context"
	"todo-api/internal/application/auth"
	"todo-api/internal/domain"
	"todo-api/internal/ports"
)

type CompleteTodo struct {
	ID int
}

type CompleteTodoHandler struct {
	mp ports.TodoMessagePublisher
}

func NewCompleteTodoHandler(mp ports.TodoMessagePublisher) *CompleteTodoHandler {
	return &CompleteTodoHandler{mp}
}

func (h *CompleteTodoHandler) Handle(ctx context.Context, c *CompleteTodo) error {
	if !auth.IsInRole(auth.GetAccessToken(ctx), "todo-user") {
		return domain.ErrUnauthorized
	}

	return h.mp.TodoComplete(ctx, c.ID)
}
