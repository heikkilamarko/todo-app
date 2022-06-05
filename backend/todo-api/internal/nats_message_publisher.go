package internal

import (
	"context"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type NATSMessagePublisher struct {
	Conn *nats.Conn
}

func (p *NATSMessagePublisher) TodoCreate(_ context.Context, todo *Todo) error {
	return p.publish("todo.create", &TodoCreateMessage{todo})
}

func (p *NATSMessagePublisher) TodoComplete(_ context.Context, id int) error {
	return p.publish("todo.complete", &TodoCompleteMessage{id})
}

func (p *NATSMessagePublisher) publish(subject string, message any) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	js, err := p.Conn.JetStream()
	if err != nil {
		return err
	}

	_, err = js.Publish(subject, data)
	return err
}
