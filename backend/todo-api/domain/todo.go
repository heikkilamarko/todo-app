// Package domain ...
package domain

import "time"

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
