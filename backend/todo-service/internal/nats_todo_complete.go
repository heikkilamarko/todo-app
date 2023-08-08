package internal

import (
	"context"
	"log/slog"

	"github.com/nats-io/nats.go/jetstream"
)

type TodoCompleteHandler struct {
	Parser *NATSMessageParser
	Repo   Repository
	Pub    MessagePublisher
	Logger *slog.Logger
}

func (h *TodoCompleteHandler) Handle(ctx context.Context, msg jetstream.Msg) error {
	_ = msg.Ack()

	message := &TodoCompleteMessage{}

	if err := h.Parser.Parse(msg, message); err != nil {
		h.Logger.Error(err.Error())
		return err
	}

	if err := h.Repo.CompleteTodo(ctx, message.ID); err != nil {
		_ = h.Pub.TodoCompleteError(ctx, "")
		return err
	}

	return h.Pub.TodoCompleteOk(ctx, message.ID)
}
