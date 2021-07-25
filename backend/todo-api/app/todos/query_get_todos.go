package todos

import (
	"net/http"
	"strconv"
	"todo-api/app/utils"

	"github.com/heikkilamarko/goutils"
)

// GetTodos method
func (c *Controller) GetTodos(w http.ResponseWriter, r *http.Request) {
	query, err := parseGetTodosRequest(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	todos, err := c.repository.getTodos(r.Context(), query)

	if err != nil {
		c.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, todos, query)
}

func parseGetTodosRequest(r *http.Request) (*getTodosQuery, error) {
	errorMap := map[string]string{}

	offset := 0
	limit := utils.LimitMaxPageSize

	var err error

	if value := r.FormValue(utils.FieldPaginationOffset); value != "" {
		offset, err = strconv.Atoi(value)
		if err != nil || offset < 0 {
			errorMap[utils.FieldPaginationOffset] = utils.ErrCodeInvalidOffset
		}
	}

	if value := r.FormValue(utils.FieldPaginationLimit); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil || limit < 1 || utils.LimitMaxPageSize < limit {
			errorMap[utils.FieldPaginationLimit] = utils.ErrCodeInvalidLimit
		}
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &getTodosQuery{offset, limit}, nil
}
