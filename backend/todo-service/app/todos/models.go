package todos

import (
	"time"
)

type todo struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type todoError struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

type createTodoCommand struct {
	Todo *todo `json:"todo"`
}

type completeTodoCommand struct {
	ID int `json:"id"`
}
