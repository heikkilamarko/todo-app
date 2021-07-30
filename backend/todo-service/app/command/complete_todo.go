// Package command ...
package command

import (
	"context"
	"todo-service/domain"
)

type CompleteTodo struct {
	ID int `json:"id"`
}

type CompleteTodoHandler struct {
	r  domain.TodoRepository
	mp domain.TodoMessagePublisher
}

func NewCompleteTodoHandler(r domain.TodoRepository, mp domain.TodoMessagePublisher) *CompleteTodoHandler {
	return &CompleteTodoHandler{r, mp}
}

func (h *CompleteTodoHandler) Handle(ctx context.Context, c *CompleteTodo) error {
	if err := h.r.CompleteTodo(ctx, c.ID); err != nil {
		_ = h.mp.TodoCompletedError(ctx, "")
		return err
	}

	if err := h.mp.TodoCompletedOk(ctx, c.ID); err != nil {
		return err
	}

	return nil
}
