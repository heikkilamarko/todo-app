package query

import (
	"context"
	"todo-api/internal/domain"
	"todo-api/internal/ports"
)

type GetTodos struct {
	Offset int
	Limit  int
}

type GetTodosHandler struct {
	r ports.TodoRepository
}

func NewGetTodosHandler(r ports.TodoRepository) *GetTodosHandler {
	return &GetTodosHandler{r}
}

func (h *GetTodosHandler) Handle(ctx context.Context, q *GetTodos) ([]*domain.Todo, error) {
	return h.r.GetTodos(ctx, &ports.GetTodosQuery{
		Offset: q.Offset,
		Limit:  q.Limit,
	})
}
