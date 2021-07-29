package command

import (
	"context"
	"todo-api/domain"
)

type CompleteTodo struct {
	ID int `json:"id"`
}

type CompleteTodoHandler struct {
	r domain.TodoRepository
}

func NewCompleteTodoHandler(r domain.TodoRepository) *CompleteTodoHandler {
	return &CompleteTodoHandler{r}
}

func (h *CompleteTodoHandler) Handle(ctx context.Context, c *CompleteTodo) error {
	return h.r.CompleteTodo(ctx, c.ID)
}
