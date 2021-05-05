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

type createTodoCommand struct {
	Todo *todo `json:"todo"`
}
