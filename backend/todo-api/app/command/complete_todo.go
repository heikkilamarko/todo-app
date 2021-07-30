package command

import (
	"context"
	"todo-api/domain"
)

type CompleteTodo struct {
	ID int `json:"id"`
}

type CompleteTodoHandler struct {
	mp domain.TodoMessagePublisher
}

func NewCompleteTodoHandler(mp domain.TodoMessagePublisher) *CompleteTodoHandler {
	return &CompleteTodoHandler{mp}
}

func (h *CompleteTodoHandler) Handle(ctx context.Context, c *CompleteTodo) error {
	return h.mp.CompleteTodo(ctx, c.ID)
}
