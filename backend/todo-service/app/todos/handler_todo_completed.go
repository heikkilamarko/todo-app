package todos

import (
	"context"
	"errors"
	"todo-service/app/utils"

	"github.com/heikkilamarko/goutils"
	"github.com/nats-io/nats.go"
)

func (c *Controller) handleTodoCompleted(ctx context.Context, m *nats.Msg) {
	m.Ack()

	command := &completeTodoCommand{}

	err := utils.
		NewMessageParser(schemaTodoCompleted).
		Parse(m.Data, command)

	if err != nil {
		c.logError(err)

		var message string

		var verr *goutils.ValidationError
		if errors.As(err, &verr) {
			message = verr.Error()
		}

		c.publishMessage(
			subjectTodoCompletedError,
			errorMessage{subjectTodoCompletedError, message},
		)

		return
	}

	if err := c.repository.completeTodo(ctx, command); err != nil {
		c.logError(err)
		c.publishMessage(
			subjectTodoCompletedError,
			errorMessage{Code: subjectTodoCompletedError},
		)
		return
	}

	if err := c.publishMessage(subjectTodoCompletedOk, command); err != nil {
		c.logError(err)
	}
}
