package workflow

import (
	"context"
	"database/sql"
)

type Activities struct {
	db *sql.DB
}

func NewActivities(db *sql.DB) *Activities {
	return &Activities{db}
}

func (a *Activities) RemoveTodos(ctx context.Context) error {
	_, err := a.db.ExecContext(ctx, "DELETE FROM todos")
	return err
}
