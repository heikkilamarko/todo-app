package internal

import (
	"context"
	"encoding/json"

	"github.com/centrifugal/gocent/v3"
)

type CentrifugoMessagePublisher struct {
	config *Config
	client *gocent.Client
}

func NewCentrifugoMessagePublisher(c *Config) *CentrifugoMessagePublisher {
	return &CentrifugoMessagePublisher{
		c,
		gocent.New(gocent.Config{Addr: c.CentrifugoURL, Key: c.CentrifugoKey}),
	}
}

func (p *CentrifugoMessagePublisher) TodoCreateOk(ctx context.Context, todo *Todo) error {
	return p.Publish(ctx,
		"todo.create.ok",
		&TodoCreateOkMessage{todo},
	)
}

func (p *CentrifugoMessagePublisher) TodoCreateError(ctx context.Context, message string) error {
	return p.Publish(ctx,
		"todo.create.error",
		&ErrorCode{"todo.create.error", message},
	)
}

func (p *CentrifugoMessagePublisher) TodoCompleteOk(ctx context.Context, id int) error {
	return p.Publish(ctx,
		"todo.complete.ok",
		&TodoCompleteOkMessage{id},
	)
}

func (p *CentrifugoMessagePublisher) TodoCompleteError(ctx context.Context, message string) error {
	return p.Publish(ctx,
		"todo.complete.error",
		&ErrorCode{"todo.complete.error", message},
	)
}

func (p *CentrifugoMessagePublisher) Publish(ctx context.Context, t string, d any) error {
	data, err := json.Marshal(MessageWrapper{t, d})
	if err != nil {
		return err
	}

	_, err = p.client.Publish(ctx, "notifications", data)
	return err
}
