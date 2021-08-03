package command

import (
	"context"
	"todo-service/internal/ports"
)

type CompleteTodo struct {
	ID int
}

type CompleteTodoHandler struct {
	r  ports.TodoRepository
	mp ports.TodoMessagePublisher
}

func NewCompleteTodoHandler(r ports.TodoRepository, mp ports.TodoMessagePublisher) *CompleteTodoHandler {
	return &CompleteTodoHandler{r, mp}
}

func (h *CompleteTodoHandler) Handle(ctx context.Context, c *CompleteTodo) error {
	if err := h.r.CompleteTodo(ctx, c.ID); err != nil {
		_ = h.mp.TodoCompleteError(ctx, "")
		return err
	}

	return h.mp.TodoCompleteOk(ctx, c.ID)
}
