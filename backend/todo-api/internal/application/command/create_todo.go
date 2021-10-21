package command

import (
	"context"
	"todo-api/internal/application/auth"
	"todo-api/internal/domain"
	"todo-api/internal/ports"
)

type CreateTodo struct {
	Todo *domain.Todo
}

type CreateTodoHandler struct {
	mp ports.TodoMessagePublisher
}

func NewCreateTodoHandler(mp ports.TodoMessagePublisher) *CreateTodoHandler {
	return &CreateTodoHandler{mp}
}

func (h *CreateTodoHandler) Handle(ctx context.Context, c *CreateTodo) error {
	if !auth.IsInRole(auth.GetAccessToken(ctx), "todo-user") {
		return domain.ErrUnauthorized
	}

	return h.mp.TodoCreate(ctx, c.Todo)
}
