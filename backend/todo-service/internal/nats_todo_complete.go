package internal

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

type TodoCompleteHandler struct {
	Parser *NATSMessageParser
	Repo   Repository
	Pub    MessagePublisher
	Logger *zerolog.Logger
}

func (h *TodoCompleteHandler) Handle(ctx context.Context, m *nats.Msg) error {
	_ = m.Ack()

	message := &TodoCompleteMessage{}

	if err := h.Parser.Parse(m, message); err != nil {
		h.Logger.Error().Err(err).Send()
		return err
	}

	if err := h.Repo.CompleteTodo(ctx, message.ID); err != nil {
		_ = h.Pub.TodoCompleteError(ctx, "")
		return err
	}

	return h.Pub.TodoCompleteOk(ctx, message.ID)
}
