package internal

import (
	"context"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

type TodoCreateHandler struct {
	parser *NATSMessageParser
	repo   Repository
	pub    MessagePublisher
	logger *zerolog.Logger
}

func (h *TodoCreateHandler) Handle(ctx context.Context, m *nats.Msg) error {
	_ = m.Ack()

	message := &TodoCreateMessage{}

	if err := h.parser.Parse(m, message); err != nil {
		h.logger.Error().Err(err).Send()
		return err
	}

	message.Todo.SetCreateTimestamps()

	if err := h.repo.CreateTodo(ctx, message.Todo); err != nil {
		_ = h.pub.TodoCreateError(ctx, "")
		return err
	}

	return h.pub.TodoCreateOk(ctx, message.Todo)
}
