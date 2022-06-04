package internal

import (
	"context"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

const (
	subjectTodoCreate   = "todo.create"
	subjectTodoComplete = "todo.complete"
)

type todoCreateMessage struct {
	Todo *Todo `json:"todo"`
}

type todoCompleteMessage struct {
	ID int `json:"id"`
}

type NATSTodoMessagePublisher struct {
	js nats.JetStreamContext
}

func NewNATSTodoMessagePublisher(js nats.JetStreamContext) *NATSTodoMessagePublisher {
	return &NATSTodoMessagePublisher{js}
}

func (mp *NATSTodoMessagePublisher) TodoCreate(_ context.Context, todo *Todo) error {
	return mp.publish(subjectTodoCreate, &todoCreateMessage{todo})
}

func (mp *NATSTodoMessagePublisher) TodoComplete(_ context.Context, id int) error {
	return mp.publish(subjectTodoComplete, &todoCompleteMessage{id})
}

func (mp *NATSTodoMessagePublisher) publish(subject string, message any) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = mp.js.Publish(subject, data)
	return err
}
