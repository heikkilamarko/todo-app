package todos

import (
	"context"
	"errors"

	"github.com/heikkilamarko/goutils"
	"github.com/nats-io/nats.go"
)

func (c *Controller) handleTodoCompleted(ctx context.Context, m *nats.Msg) {
	_ = m.Ack()

	command := &completeTodoCommand{}

	if err := c.messageParser.Parse(m, command); err != nil {
		c.logError(err)

		var message string

		var verr *goutils.ValidationError
		if errors.As(err, &verr) {
			message = verr.Error()
		}

		_ = c.publishMessage(
			subjectTodoCompletedError,
			errorMessage{
				Code:    subjectTodoCompletedError,
				Message: message,
			},
		)

		return
	}

	if err := c.repository.completeTodo(ctx, command); err != nil {
		c.logError(err)

		_ = c.publishMessage(
			subjectTodoCompletedError,
			errorMessage{
				Code: subjectTodoCompletedError,
			},
		)

		return
	}

	if err := c.publishMessage(subjectTodoCompletedOk, command); err != nil {
		c.logError(err)
	}
}
