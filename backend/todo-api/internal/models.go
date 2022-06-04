package internal

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type PaginationMeta struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type TodoCreateMessage struct {
	Todo *Todo `json:"todo"`
}

type TodoCompleteMessage struct {
	ID int `json:"id"`
}
