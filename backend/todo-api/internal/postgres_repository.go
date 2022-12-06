package internal

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"strings"

	"github.com/samber/lo"
)

//go:embed sql/get_permissions.sql
var getPermissionsSQL string

//go:embed sql/get_todos.sql
var getTodosSQL string

type PostgresRepository struct {
	DB *sql.DB
}

func (r *PostgresRepository) GetPermissions(ctx context.Context, roles []string) ([]string, error) {
	n := len(roles)

	if n == 0 {
		return []string{}, nil
	}

	query := strings.Replace(getPermissionsSQL, "$1", buildPostgresInParams(n), 1)
	args := lo.ToAnySlice(roles)

	rows, err := r.DB.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := []string{}

	for rows.Next() {
		var d string

		err := rows.Scan(&d)

		if err != nil {
			return nil, err
		}

		data = append(data, d)
	}

	return data, nil
}

func (r *PostgresRepository) GetTodos(ctx context.Context, q *GetTodosQuery) ([]*Todo, error) {
	rows, err := r.DB.QueryContext(
		ctx,
		getTodosSQL,
		q.Limit, q.Offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := []*Todo{}

	for rows.Next() {
		d := &Todo{}

		err := rows.Scan(
			&d.ID,
			&d.Name,
			&d.Description,
			&d.CreatedAt,
			&d.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		data = append(data, d)
	}

	return data, nil
}

func buildPostgresInParams(n int) string {
	return strings.Join(
		lo.RepeatBy(n, func(i int) string {
			return fmt.Sprintf("$%d", i+1)
		}),
		",")
}
