package internal

import (
	"context"
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type NATSMessagePublisher struct {
	Conn *nats.Conn
}

func (p *NATSMessagePublisher) TodoCreate(ctx context.Context, todo *Todo) error {
	return p.publish(ctx, "todo.create", &TodoCreateMessage{todo})
}

func (p *NATSMessagePublisher) TodoComplete(ctx context.Context, id int) error {
	return p.publish(ctx, "todo.complete", &TodoCompleteMessage{id})
}

func (p *NATSMessagePublisher) publish(ctx context.Context, subject string, message any) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	js, err := jetstream.New(p.Conn)
	if err != nil {
		return err
	}

	_, err = js.Publish(ctx, subject, data)
	return err
}
