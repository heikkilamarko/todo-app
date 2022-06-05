package internal

import (
	"context"
	"database/sql"
)

type Activities struct {
	DB *sql.DB
}

func (a *Activities) RemoveTodos(ctx context.Context) error {
	_, err := a.DB.ExecContext(ctx, "DELETE FROM todos")
	return err
}
