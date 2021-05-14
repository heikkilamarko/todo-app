package todos

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/app/utils"

	"github.com/gorilla/mux"
)

// CompleteTodo method
func (c *Controller) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	command, err := parseCompleteTodoRequest(r)

	if err != nil {
		utils.WriteValidationError(w, err)
		return
	}

	data, err := json.Marshal(command)

	if err != nil {
		c.logger.Error().Err(err).Send()
		return
	}

	_, err = c.js.Publish(subjectTodoCompleted, data)

	if err != nil {
		c.logger.Error().Err(err).Send()
		utils.WriteInternalError(w, nil)
		return
	}

	utils.WriteResponse(w, http.StatusOK, nil)
}

func parseCompleteTodoRequest(r *http.Request) (*completeTodoCommand, error) {
	validationErrors := map[string]string{}

	id, err := strconv.Atoi(mux.Vars(r)[utils.FieldID])
	if err != nil {
		validationErrors[utils.FieldID] = utils.ErrCodeInvalidID
	}

	if 0 < len(validationErrors) {
		return nil, utils.NewValidationError(validationErrors)
	}

	return &completeTodoCommand{id}, nil
}
