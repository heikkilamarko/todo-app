package todos

import (
	"encoding/json"
	"net/http"
	"todo-api/app/utils"
)

// CreateTodo method
func (c *Controller) CreateTodo(w http.ResponseWriter, r *http.Request) {
	command, err := parseCreateTodoRequest(r)

	if err != nil {
		utils.WriteValidationError(w, err)
		return
	}

	data, err := json.Marshal(command)

	if err != nil {
		c.logger.Error().Err(err).Send()
		return
	}

	_, err = c.js.Publish(subjectTodoCreated, data)

	if err != nil {
		c.logger.Error().Err(err).Send()
		utils.WriteInternalError(w, nil)
		return
	}

	utils.WriteResponse(w, http.StatusOK, nil)
}

func parseCreateTodoRequest(r *http.Request) (*createTodoCommand, error) {
	validationErrors := map[string]string{}

	todo := &todo{}
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		validationErrors[utils.FieldRequestBody] = utils.ErrCodeInvalidRequestBody
	}

	if 0 < len(validationErrors) {
		return nil, utils.NewValidationError(validationErrors)
	}

	return &createTodoCommand{todo}, nil
}
