package todos

import (
	"net/http"
	"strconv"
	"todo-api/app/utils"

	"github.com/gorilla/mux"
)

// CompleteTodo method
func (c *Controller) CompleteTodo(w http.ResponseWriter, r *http.Request) {
	command, err := parseCompleteTodoRequest(r)

	if err != nil {
		c.logError(err)
		utils.WriteValidationError(w, err)
		return
	}

	if err := c.publishMessage(subjectTodoCompleted, command); err != nil {
		c.logError(err)
		utils.WriteInternalError(w, nil)
		return
	}

	utils.WriteResponse(w, http.StatusOK, nil)
}

func parseCompleteTodoRequest(r *http.Request) (*completeTodoCommand, error) {
	errorMap := map[string]string{}

	id, err := strconv.Atoi(mux.Vars(r)[utils.FieldID])
	if err != nil {
		errorMap[utils.FieldID] = utils.ErrCodeInvalidID
	}

	if 0 < len(errorMap) {
		return nil, utils.NewValidationError(errorMap)
	}

	return &completeTodoCommand{id}, nil
}
