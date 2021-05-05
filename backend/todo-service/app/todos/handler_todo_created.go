package todos

import (
	"context"
	_ "embed"
	"encoding/json"
	"todo-service/app/utils"

	"github.com/nats-io/nats.go"
)

const subjectTodoCreated = "todo.created"
const subjectTodoProcessed = "todo.processed"

//go:embed schemas/new-todo.json
var newTodoSchema string

func (c *Controller) handleTodoCreated() error {

	validator := utils.NewJSONSchemaValidator(newTodoSchema)

	_, err := c.natsConn.Subscribe(subjectTodoCreated, func(m *nats.Msg) {

		c.logger.Info().Msgf("Message received (%s)", m.Subject)

		ves, err := validator.Validate(string(m.Data))

		if err != nil {
			c.logger.Error().Err(err).Send()
			return
		}

		if 0 < len(ves) {
			c.logger.Warn().Msg("Invalid message")
			for _, ve := range ves {
				c.logger.Warn().Msgf("Validation error: %s", ve)
			}
			return
		}

		todo := &todo{}

		if err := json.Unmarshal(m.Data, todo); err != nil {
			c.logger.Error().Err(err).Send()
			return
		}

		command := &createTodoCommand{Todo: todo}

		if err := c.repository.createTodo(context.Background(), command); err != nil {
			c.logger.Error().Err(err).Send()
			return
		}

		todoBytes, err := json.Marshal(todo)

		if err != nil {
			c.logger.Error().Err(err).Send()
			return
		}

		if err := c.natsConn.Publish(subjectTodoProcessed, todoBytes); err != nil {
			c.logger.Error().Err(err).Send()
			return
		}
	})

	if err != nil {
		return err
	}

	return nil
}
