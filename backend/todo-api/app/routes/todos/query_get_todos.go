package todos

import (
	"net/http"
	"strconv"
	"todo-api/app/constants"

	"github.com/heikkilamarko/goutils"
)

// GetTodos query
func (c *Controller) GetTodos(w http.ResponseWriter, r *http.Request) {
	query, err := parseGetTodosRequest(r)

	if err != nil {
		goutils.WriteValidationError(w, err)
		return
	}

	todos, err := c.Repository.GetTodos(r.Context(), query)

	if err != nil {
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, todos, query)
}

func parseGetTodosRequest(r *http.Request) (*GetTodosQuery, error) {
	validationErrors := map[string]string{}

	var offset int = 0
	var limit int = constants.PaginationLimitMax

	var err error = nil

	if value := r.FormValue(constants.FieldPaginationOffset); value != "" {
		offset, err = strconv.Atoi(value)
		if err != nil || offset < 0 {
			validationErrors[constants.FieldPaginationOffset] = constants.ErrCodeInvalidOffset
		}
	}

	if value := r.FormValue(constants.FieldPaginationLimit); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil || limit < 1 || constants.PaginationLimitMax < limit {
			validationErrors[constants.FieldPaginationLimit] = constants.ErrCodeInvalidLimit
		}
	}

	if 0 < len(validationErrors) {
		return nil, goutils.NewValidationError(validationErrors)
	}

	return &GetTodosQuery{offset, limit}, nil
}
