package command

import (
	"context"
	"todo-api/internal/ports"
)

type CompleteTodo struct {
	ID int `json:"id"`
}

type CompleteTodoHandler struct {
	mp ports.TodoMessagePublisher
}

func NewCompleteTodoHandler(mp ports.TodoMessagePublisher) *CompleteTodoHandler {
	return &CompleteTodoHandler{mp}
}

func (h *CompleteTodoHandler) Handle(ctx context.Context, c *CompleteTodo) error {
	return h.mp.TodoComplete(ctx, c.ID)
}
