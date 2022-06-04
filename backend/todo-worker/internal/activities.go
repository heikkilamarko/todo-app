package internal

import (
	"context"
	"database/sql"
)

type Activities struct {
	db *sql.DB
}

func (a *Activities) RemoveTodos(ctx context.Context) error {
	_, err := a.db.ExecContext(ctx, "DELETE FROM todos")
	return err
}
