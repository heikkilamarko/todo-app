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
		c.logError(err)
		utils.WriteValidationError(w, err)
		return
	}

	if err := c.publishMessage(subjectTodoCreated, command); err != nil {
		c.logError(err)
		utils.WriteInternalError(w, nil)
		return
	}

	utils.WriteResponse(w, http.StatusAccepted, nil)
}

func parseCreateTodoRequest(r *http.Request) (*createTodoCommand, error) {
	errorMap := map[string]string{}

	todo := &todo{}
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		errorMap[utils.FieldRequestBody] = utils.ErrCodeInvalidRequestBody
	}

	if 0 < len(errorMap) {
		return nil, utils.NewValidationError(errorMap)
	}

	return &createTodoCommand{todo}, nil
}
