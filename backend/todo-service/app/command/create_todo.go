package command

import (
	"context"
	"todo-service/domain"
)

type CreateTodo struct {
	Todo *domain.Todo `json:"todo"`
}

type CreateTodoHandler struct {
	r  domain.TodoRepository
	mp domain.TodoMessagePublisher
}

func NewCreateTodoHandler(r domain.TodoRepository, mp domain.TodoMessagePublisher) *CreateTodoHandler {
	return &CreateTodoHandler{r, mp}
}

func (h *CreateTodoHandler) Handle(ctx context.Context, c *CreateTodo) error {
	if err := h.r.CreateTodo(ctx, c.Todo); err != nil {
		_ = h.mp.TodoCreatedError(ctx, "")
		return err
	}

	if err := h.mp.TodoCreatedOk(ctx, c.Todo); err != nil {
		return err
	}

	return nil
}
