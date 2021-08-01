package adapters

import (
	"context"
	"encoding/json"
	"todo-service/internal/domain"

	"github.com/nats-io/nats.go"
)

const (
	subjectTodoCreateOk      = "todo.create.ok"
	subjectTodoCreateError   = "todo.create.error"
	subjectTodoCompleteOk    = "todo.complete.ok"
	subjectTodoCompleteError = "todo.complete.error"
)

type todoCreateOkMessage struct {
	Todo *domain.Todo `json:"todo"`
}

type todoCompleteOkMessage struct {
	ID int `json:"id"`
}

type TodoMessagePublisher struct {
	nc *nats.Conn
}

func NewTodoMessagePublisher(nc *nats.Conn) *TodoMessagePublisher {
	return &TodoMessagePublisher{nc}
}

func (mp *TodoMessagePublisher) TodoCreateOk(_ context.Context, todo *domain.Todo) error {
	return mp.publish(subjectTodoCreateOk, &todoCreateOkMessage{todo})
}

func (mp *TodoMessagePublisher) TodoCreateError(_ context.Context, message string) error {
	return mp.publish(
		subjectTodoCreateError,
		&domain.ErrorMessage{
			Code:    subjectTodoCreateError,
			Message: message,
		},
	)
}

func (mp *TodoMessagePublisher) TodoCompleteOk(_ context.Context, id int) error {
	return mp.publish(subjectTodoCompleteOk, &todoCompleteOkMessage{id})
}

func (mp *TodoMessagePublisher) TodoCompleteError(_ context.Context, message string) error {
	return mp.publish(
		subjectTodoCompleteError,
		&domain.ErrorMessage{
			Code:    subjectTodoCompleteError,
			Message: message,
		},
	)
}

func (mp *TodoMessagePublisher) publish(subject string, message interface{}) error {
	data, err := json.Marshal(message)

	if err != nil {
		return err
	}

	if err := mp.nc.Publish(subject, data); err != nil {
		return err
	}

	return nil
}
