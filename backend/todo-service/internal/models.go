package internal

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

func (t *Todo) SetCreateTimestamps() {
	now := time.Now()
	t.CreatedAt = now
	t.UpdatedAt = now
}

type TodoCreateMessage struct {
	Todo *Todo `json:"todo"`
}

type TodoCreateOkMessage struct {
	Todo *Todo `json:"todo"`
}

type TodoCompleteMessage struct {
	ID int `json:"id"`
}

type TodoCompleteOkMessage struct {
	ID int `json:"id"`
}

type MessageWrapper struct {
	Type string `json:"type"`
	Data any    `json:"data"`
}
