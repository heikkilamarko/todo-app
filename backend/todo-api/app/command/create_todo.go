// Package command ...
package command

import (
	"context"
	"todo-api/domain"
)

type CreateTodo struct {
	Todo *domain.Todo `json:"todo"`
}

type CreateTodoHandler struct {
	r domain.TodoRepository
}

func NewCreateTodoHandler(r domain.TodoRepository) *CreateTodoHandler {
	return &CreateTodoHandler{r}
}

func (h *CreateTodoHandler) Handle(ctx context.Context, c *CreateTodo) error {
	return h.r.CreateTodo(ctx, c.Todo)
}
