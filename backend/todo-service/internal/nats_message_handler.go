package internal

import (
	"context"

	"github.com/nats-io/nats.go/jetstream"
)

type NATSMessageHandler interface {
	Handle(ctx context.Context, msg jetstream.Msg) error
}
