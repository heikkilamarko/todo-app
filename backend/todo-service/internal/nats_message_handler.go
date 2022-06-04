package internal

import (
	"context"

	"github.com/nats-io/nats.go"
)

type NATSMessageHandler interface {
	Handle(ctx context.Context, m *nats.Msg) error
}
