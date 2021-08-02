package command

import (
	"context"
	"todo-service/internal/domain"
	"todo-service/internal/ports"
)

type CreateTodo struct {
	Todo *domain.Todo `json:"todo"`
}

type CreateTodoHandler struct {
	r  ports.TodoRepository
	mp ports.TodoMessagePublisher
}

func NewCreateTodoHandler(r ports.TodoRepository, mp ports.TodoMessagePublisher) *CreateTodoHandler {
	return &CreateTodoHandler{r, mp}
}

func (h *CreateTodoHandler) Handle(ctx context.Context, c *CreateTodo) error {
	if err := h.r.CreateTodo(ctx, c.Todo); err != nil {
		_ = h.mp.TodoCreateError(ctx, "")
		return err
	}

	return h.mp.TodoCreateOk(ctx, c.Todo)
}
