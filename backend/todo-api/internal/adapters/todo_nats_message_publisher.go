package adapters

import (
	"context"
	"encoding/json"
	"todo-api/internal/domain"

	"github.com/nats-io/nats.go"
)

const (
	subjectTodoCreate   = "todo.create"
	subjectTodoComplete = "todo.complete"
)

type todoCreateMessage struct {
	Todo *domain.Todo `json:"todo"`
}

type todoCompleteMessage struct {
	ID int `json:"id"`
}

type TodoNATSMessagePublisher struct {
	js nats.JetStreamContext
}

func NewTodoNATSMessagePublisher(js nats.JetStreamContext) *TodoNATSMessagePublisher {
	return &TodoNATSMessagePublisher{js}
}

func (mp *TodoNATSMessagePublisher) TodoCreate(_ context.Context, todo *domain.Todo) error {
	return mp.publish(subjectTodoCreate, &todoCreateMessage{todo})
}

func (mp *TodoNATSMessagePublisher) TodoComplete(_ context.Context, id int) error {
	return mp.publish(subjectTodoComplete, &todoCompleteMessage{id})
}

func (mp *TodoNATSMessagePublisher) publish(subject string, message interface{}) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = mp.js.Publish(subject, data)
	return err
}
