package internal

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

type TodoCompleteHandler struct {
	parser *NATSMessageParser
	repo   Repository
	pub    MessagePublisher
	logger *zerolog.Logger
}

func (h *TodoCompleteHandler) Handle(ctx context.Context, m *nats.Msg) error {
	_ = m.Ack()

	message := &TodoCompleteMessage{}

	if err := h.parser.Parse(m, message); err != nil {
		h.logger.Error().Err(err).Send()
		return err
	}

	if err := h.repo.CompleteTodo(ctx, message.ID); err != nil {
		_ = h.pub.TodoCompleteError(ctx, "")
		return err
	}

	return h.pub.TodoCompleteOk(ctx, message.ID)
}
