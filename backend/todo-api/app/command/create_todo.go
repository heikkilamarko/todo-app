package command

import (
	"context"
	"todo-api/domain"
)

type CreateTodo struct {
	Todo *domain.Todo `json:"todo"`
}

type CreateTodoHandler struct {
	mp domain.TodoMessagePublisher
}

func NewCreateTodoHandler(mp domain.TodoMessagePublisher) *CreateTodoHandler {
	return &CreateTodoHandler{mp}
}

func (h *CreateTodoHandler) Handle(ctx context.Context, c *CreateTodo) error {
	return h.mp.TodoCreate(ctx, c.Todo)
}
