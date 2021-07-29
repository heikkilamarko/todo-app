// Package adapters ...
package adapters

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"todo-api/domain"

	"github.com/nats-io/nats.go"
)

//go:embed sql/get_todos.sql
var getTodosSQL string

const (
	subjectTodoCreated   = "todo.created"
	subjectTodoCompleted = "todo.completed"
)

type createTodoMessage struct {
	Todo *domain.Todo `json:"todo"`
}

type completeTodoMessage struct {
	ID int `json:"id"`
}

type TodoRepository struct {
	db *sql.DB
	js nats.JetStreamContext
}

func NewTodoRepository(db *sql.DB, js nats.JetStreamContext) *TodoRepository {
	return &TodoRepository{db, js}
}

func (r *TodoRepository) GetTodos(ctx context.Context, query *domain.GetTodosQuery) ([]*domain.Todo, error) {
	rows, err := r.db.QueryContext(
		ctx,
		getTodosSQL,
		query.Limit, query.Offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []*domain.Todo{}

	for rows.Next() {
		t := &domain.Todo{}

		err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.Description,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, t)
	}

	return todos, nil
}

func (r *TodoRepository) CreateTodo(_ context.Context, todo *domain.Todo) error {
	return r.publish(subjectTodoCreated, &createTodoMessage{todo})
}

func (r *TodoRepository) CompleteTodo(_ context.Context, id int) error {
	return r.publish(subjectTodoCompleted, &completeTodoMessage{id})
}

func (r *TodoRepository) publish(subject string, message interface{}) error {
	data, err := json.Marshal(message)

	if err != nil {
		return err
	}

	if _, err := r.js.Publish(subject, data); err != nil {
		return err
	}

	return nil
}
