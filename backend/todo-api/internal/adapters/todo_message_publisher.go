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

type TodoMessagePublisher struct {
	js nats.JetStreamContext
}

func NewTodoMessagePublisher(js nats.JetStreamContext) *TodoMessagePublisher {
	return &TodoMessagePublisher{js}
}

func (mp *TodoMessagePublisher) TodoCreate(_ context.Context, todo *domain.Todo) error {
	return mp.publish(subjectTodoCreate, &todoCreateMessage{todo})
}

func (mp *TodoMessagePublisher) TodoComplete(_ context.Context, id int) error {
	return mp.publish(subjectTodoComplete, &todoCompleteMessage{id})
}

func (mp *TodoMessagePublisher) publish(subject string, message interface{}) error {
	data, err := json.Marshal(message)

	if err != nil {
		return err
	}

	if _, err := mp.js.Publish(subject, data); err != nil {
		return err
	}

	return nil
}
