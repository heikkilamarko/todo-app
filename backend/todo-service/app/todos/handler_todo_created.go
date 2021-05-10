package todos

import (
	"context"
	_ "embed"
	"encoding/json"
	"todo-service/app/utils"
)

//go:embed schemas/new-todo.json
var newTodoSchema string

func (c *Controller) handleTodoCreated() {

	go func() {

		sub, err := c.js.PullSubscribe("todo.created", "TODOS")

		if err != nil {
			c.logger.Error().Err(err).Send()
			return
		}

		validator := utils.NewJSONSchemaValidator(newTodoSchema)

		for {
			ms, err := sub.Fetch(1)
			if err != nil {
				continue
			}

			m := ms[0]

			c.logger.Info().Msgf("Message received (%s)", m.Subject)

			ves, err := validator.Validate(string(m.Data))

			if err != nil {
				c.logger.Error().Err(err).Send()
				continue
			}

			if 0 < len(ves) {
				c.logger.Warn().Msg("Invalid message")
				for _, ve := range ves {
					c.logger.Warn().Msgf("Validation error: %s", ve)
				}
				continue
			}

			todo := &todo{}

			if err := json.Unmarshal(m.Data, todo); err != nil {
				c.logger.Error().Err(err).Send()
				continue
			}

			command := &createTodoCommand{Todo: todo}

			if err := c.repository.createTodo(context.Background(), command); err != nil {
				c.logger.Error().Err(err).Send()
				continue
			}

			todoBytes, err := json.Marshal(todo)

			if err != nil {
				c.logger.Error().Err(err).Send()
				continue
			}

			if err := c.nc.Publish("todo.processed", todoBytes); err != nil {
				c.logger.Error().Err(err).Send()
				continue
			}

			m.Ack()
		}
	}()
}
