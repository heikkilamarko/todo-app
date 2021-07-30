package adapters

import (
	"context"
	"encoding/json"
	"todo-api/domain"

	"github.com/nats-io/nats.go"
)

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

type TodoMessagePublisher struct {
	js nats.JetStreamContext
}

func NewTodoMessagePublisher(js nats.JetStreamContext) *TodoMessagePublisher {
	return &TodoMessagePublisher{js}
}

func (mp *TodoMessagePublisher) CreateTodo(_ context.Context, todo *domain.Todo) error {
	return mp.publish(subjectTodoCreated, &createTodoMessage{todo})
}

func (mp *TodoMessagePublisher) CompleteTodo(_ context.Context, id int) error {
	return mp.publish(subjectTodoCompleted, &completeTodoMessage{id})
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
