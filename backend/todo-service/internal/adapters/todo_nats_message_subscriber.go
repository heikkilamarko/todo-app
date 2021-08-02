package adapters

import (
	"context"
	"embed"
	"todo-service/internal/adapters/utils"
	"todo-service/internal/application"
	"todo-service/internal/application/command"

	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

//go:embed schemas/*.json
var schemaFS embed.FS

const (
	subjectTodo         = "todo.*"
	durableTodo         = "todo"
	subjectTodoCreate   = "todo.create"
	subjectTodoComplete = "todo.complete"
)

type messageHandler func(context.Context, *nats.Msg)

type TodoNATSMessageSubscriber struct {
	app           *application.Application
	js            nats.JetStreamContext
	logger        *zerolog.Logger
	messageParser *utils.MessageParser
	handlerMap    map[string]messageHandler
}

func NewTodoNATSMessageSubscriber(app *application.Application, js nats.JetStreamContext, logger *zerolog.Logger) *TodoNATSMessageSubscriber {
	ms := &TodoNATSMessageSubscriber{
		app:    app,
		js:     js,
		logger: logger,
	}

	ms.messageParser = utils.NewMessageParser(utils.NewSchemaValidator(schemaFS))

	ms.handlerMap = map[string]messageHandler{
		subjectTodoCreate:   ms.todoCreate,
		subjectTodoComplete: ms.todoComplete,
	}

	return ms
}

func (ms *TodoNATSMessageSubscriber) Subscribe(ctx context.Context) error {
	sub, err := ms.js.PullSubscribe(subjectTodo, durableTodo)

	if err != nil {
		return err
	}

	go func() {
		ms.logInfo("todo message subscriber started")

		for {
			select {
			case <-ctx.Done():
				ms.logInfo("todo message subscriber stopped")
				return
			default:
			}

			messages, err := sub.Fetch(1)
			if err != nil {
				continue
			}

			message := messages[0]

			ms.logInfo("message received (%s)", message.Subject)

			handler, ok := ms.handlerMap[message.Subject]
			if ok {
				handler(ctx, message)
			} else {
				ms.logInfo("no handler found for subject '%s'", message.Subject)
			}

			ms.logInfo("message handled (%s)", message.Subject)
		}
	}()

	return nil
}

// Handlers

func (ms *TodoNATSMessageSubscriber) todoCreate(ctx context.Context, m *nats.Msg) {
	_ = m.Ack()

	command := &command.CreateTodo{}

	if err := ms.messageParser.Parse(m, command); err != nil {
		ms.logError(err)
		return
	}

	if err := ms.app.Commands.CreateTodo.Handle(ctx, command); err != nil {
		ms.logError(err)
		return
	}
}

func (ms *TodoNATSMessageSubscriber) todoComplete(ctx context.Context, m *nats.Msg) {
	_ = m.Ack()

	command := &command.CompleteTodo{}

	if err := ms.messageParser.Parse(m, command); err != nil {
		ms.logError(err)
		return
	}

	if err := ms.app.Commands.CompleteTodo.Handle(ctx, command); err != nil {
		ms.logError(err)
		return
	}
}

// Utils

func (ms *TodoNATSMessageSubscriber) logInfo(msg string, v ...interface{}) {
	ms.logger.Info().Msgf(msg, v...)
}

func (ms *TodoNATSMessageSubscriber) logError(err error) {
	ms.logger.Error().Err(err).Send()
}