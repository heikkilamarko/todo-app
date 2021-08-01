package query

import (
	"context"
	"todo-api/internal/domain"
)

type GetTodos struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetTodosHandler struct {
	r domain.TodoRepository
}

func NewGetTodosHandler(r domain.TodoRepository) *GetTodosHandler {
	return &GetTodosHandler{r}
}

func (h *GetTodosHandler) Handle(ctx context.Context, q *GetTodos) ([]*domain.Todo, error) {
	return h.r.GetTodos(ctx, &domain.GetTodosQuery{
		Offset: q.Offset,
		Limit:  q.Limit,
	})
}
