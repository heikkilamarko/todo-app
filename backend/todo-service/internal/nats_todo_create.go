package internal

import (
	"context"
	"log/slog"

	"github.com/nats-io/nats.go/jetstream"
)

type TodoCreateHandler struct {
	Parser *NATSMessageParser
	Repo   Repository
	Pub    MessagePublisher
	Logger *slog.Logger
}

func (h *TodoCreateHandler) Handle(ctx context.Context, msg jetstream.Msg) error {
	_ = msg.Ack()

	message := &TodoCreateMessage{}

	if err := h.Parser.Parse(msg, message); err != nil {
		h.Logger.Error(err.Error())
		return err
	}

	message.Todo.SetCreateTimestamps()

	if err := h.Repo.CreateTodo(ctx, message.Todo); err != nil {
		_ = h.Pub.TodoCreateError(ctx, "")
		return err
	}

	return h.Pub.TodoCreateOk(ctx, message.Todo)
}
