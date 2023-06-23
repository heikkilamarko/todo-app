package internal

import (
	"context"
	"log/slog"

	"github.com/nats-io/nats.go"
)

type TodoCreateHandler struct {
	Parser *NATSMessageParser
	Repo   Repository
	Pub    MessagePublisher
	Logger *slog.Logger
}

func (h *TodoCreateHandler) Handle(ctx context.Context, m *nats.Msg) error {
	_ = m.Ack()

	message := &TodoCreateMessage{}

	if err := h.Parser.Parse(m, message); err != nil {
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
