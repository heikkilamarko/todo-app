package internal

import "time"

type Userinfo struct {
	Permissions []string `json:"permissions"`
}

type Todo struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type GetTodosQuery struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type TodoCreateMessage struct {
	Todo *Todo `json:"todo"`
}

type TodoCompleteMessage struct {
	ID int `json:"id"`
}
