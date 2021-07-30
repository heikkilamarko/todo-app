package adapters

import (
	"context"
	"encoding/json"
	"todo-service/domain"

	"github.com/nats-io/nats.go"
)

const (
	subjectTodoCreatedOk      = "todo.created.ok"
	subjectTodoCreatedError   = "todo.created.error"
	subjectTodoCompletedOk    = "todo.completed.ok"
	subjectTodoCompletedError = "todo.completed.error"
)

type todoCreatedOkMessage struct {
	Todo *domain.Todo `json:"todo"`
}

type todoCompletedOkMessage struct {
	ID int `json:"id"`
}

type TodoMessagePublisher struct {
	nc *nats.Conn
}

func NewTodoMessagePublisher(nc *nats.Conn) *TodoMessagePublisher {
	return &TodoMessagePublisher{nc}
}

func (mp *TodoMessagePublisher) TodoCreatedOk(_ context.Context, todo *domain.Todo) error {
	return mp.publish(subjectTodoCreatedOk, &todoCreatedOkMessage{todo})
}

func (mp *TodoMessagePublisher) TodoCreatedError(_ context.Context, message string) error {
	return mp.publish(
		subjectTodoCreatedError,
		&domain.ErrorMessage{
			Code:    subjectTodoCreatedError,
			Message: message,
		},
	)
}

func (mp *TodoMessagePublisher) TodoCompletedOk(_ context.Context, id int) error {
	return mp.publish(subjectTodoCompletedOk, &todoCompletedOkMessage{id})
}

func (mp *TodoMessagePublisher) TodoCompletedError(_ context.Context, message string) error {
	return mp.publish(
		subjectTodoCompletedError,
		&domain.ErrorMessage{
			Code:    subjectTodoCompletedError,
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
