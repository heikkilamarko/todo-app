package todos

import (
	"context"
	"errors"
	"todo-service/app/utils"

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

		var verr *utils.ValidationError
		if errors.As(err, &verr) {
			message = verr.Error()
		}

		c.publishMessage(
			subjectTodoCompletedError,
			utils.ErrorMessage{
				Code:    subjectTodoCompletedError,
				Message: message,
			},
		)

		return
	}

	if err := c.repository.completeTodo(ctx, command); err != nil {
		c.logError(err)
		c.publishMessage(
			subjectTodoCompletedError,
			utils.ErrorMessage{
				Code: subjectTodoCompletedError,
			},
		)
		return
	}

	if err := c.publishMessage(subjectTodoCompletedOk, command); err != nil {
		c.logError(err)
	}
}
