package todos

import (
	"context"
	_ "embed"
	"encoding/json"
	"todo-service/app/constants"
	"todo-service/app/utils"

	"github.com/nats-io/nats.go"
)

//go:embed schemas/todo.created.json
var todoCreatedSchema string

func (c *Controller) handleTodoCreated(ctx context.Context, m *nats.Msg) {
	m.Ack()

	validator := utils.NewJSONSchemaValidator(todoCreatedSchema)

	ves, err := validator.Validate(string(m.Data))

	if err != nil {
		c.logError(err)
		c.publishError(constants.MessageTodoCreatedError, "")
		return
	}

	if 0 < len(ves) {
		c.logger.Warn().Msg("invalid message")
		for _, ve := range ves {
			c.logger.Warn().Msgf("validation error: %s", ve)
		}
		c.publishError(constants.MessageTodoCreatedError, "bad request")
		return
	}

	todo := &todo{}

	if err := json.Unmarshal(m.Data, todo); err != nil {
		c.logError(err)
		c.publishError(constants.MessageTodoCreatedError, "")
		return
	}

	command := &createTodoCommand{Todo: todo}

	if err := c.repository.createTodo(ctx, command); err != nil {
		c.logError(err)
		c.publishError(constants.MessageTodoCreatedError, "")
		return
	}

	data, err := json.Marshal(todo)

	if err != nil {
		c.logError(err)
		return
	}

	if err := c.nc.Publish(constants.MessageTodoCreatedOk, data); err != nil {
		c.logError(err)
		return
	}
}
