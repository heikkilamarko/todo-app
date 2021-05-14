package todos

import (
	_ "embed"
	"time"
)

const (
	subjectTodoCreated   = "todo.created"
	subjectTodoCompleted = "todo.completed"
)

var (
	//go:embed sql/get_todos.sql
	sqlGetTodos string
)

type todo struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type getTodosQuery struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type createTodoCommand struct {
	Todo *todo `json:"todo"`
}

type completeTodoCommand struct {
	ID int `json:"id"`
}