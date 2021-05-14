package todos

import (
	_ "embed"
	"time"
)

const (
	subjectTodo               = "todo.*"
	subjectTodoCreated        = "todo.created"
	subjectTodoCreatedOk      = "todo.created.ok"
	subjectTodoCreatedError   = "todo.created.error"
	subjectTodoCompleted      = "todo.completed"
	subjectTodoCompletedOk    = "todo.completed.ok"
	subjectTodoCompletedError = "todo.completed.error"
	durableTodo               = "todo"
)

var (
	//go:embed schemas/todo.created.json
	schemaTodoCreated string
	//go:embed schemas/todo.completed.json
	schemaTodoCompleted string
	//go:embed sql/create_todo.sql
	sqlCreateTodo string
	//go:embed sql/complete_todo.sql
	sqlCompleteTodo string
)

type todo struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type createTodoCommand struct {
	Todo *todo `json:"todo"`
}

type completeTodoCommand struct {
	ID int `json:"id"`
}

type errorMessage struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}
