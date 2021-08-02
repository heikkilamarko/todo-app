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

type errorMessage struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

type TodoNATSMessagePublisher struct {
	nc *nats.Conn
}

func NewTodoNATSMessagePublisher(nc *nats.Conn) *TodoNATSMessagePublisher {
	return &TodoNATSMessagePublisher{nc}
}

func (mp *TodoNATSMessagePublisher) TodoCreateOk(_ context.Context, todo *domain.Todo) error {
	return mp.publish(subjectTodoCreateOk, &todoCreateOkMessage{todo})
}

func (mp *TodoNATSMessagePublisher) TodoCreateError(_ context.Context, message string) error {
	return mp.publish(
		subjectTodoCreateError,
		&errorMessage{
			Code:    subjectTodoCreateError,
			Message: message,
		},
	)
}

func (mp *TodoNATSMessagePublisher) TodoCompleteOk(_ context.Context, id int) error {
	return mp.publish(subjectTodoCompleteOk, &todoCompleteOkMessage{id})
}

func (mp *TodoNATSMessagePublisher) TodoCompleteError(_ context.Context, message string) error {
	return mp.publish(
		subjectTodoCompleteError,
		&errorMessage{
			Code:    subjectTodoCompleteError,
			Message: message,
		},
	)
}

func (mp *TodoNATSMessagePublisher) publish(subject string, message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return mp.nc.Publish(subject, data)
}
