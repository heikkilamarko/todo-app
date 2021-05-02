package todos

import (
	"context"
	"time"
)

// Todo struct
type Todo struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// CreateTodoCommand struct
type CreateTodoCommand struct {
	Todo *Todo `json:"todo"`
}

// Repository interface
type Repository interface {
	CreateTodo(context.Context, *CreateTodoCommand) error
}
