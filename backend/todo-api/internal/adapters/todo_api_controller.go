package adapters

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo-api/internal/app"
	"todo-api/internal/app/command"
	"todo-api/internal/app/query"
	"todo-api/internal/domain"

	"github.com/gorilla/mux"
	"github.com/heikkilamarko/goutils"
	"github.com/rs/zerolog"
)

const (
	errCodeInvalidID          = "invalid_id"
	errCodeInvalidOffset      = "invalid_offset"
	errCodeInvalidLimit       = "invalid_limit"
	errCodeInvalidRequestBody = "invalid_request_body"
)

const (
	fieldID               = "id"
	fieldPaginationOffset = "offset"
	fieldPaginationLimit  = "limit"
	fieldRequestBody      = "request_body"
)

const (
	limitMaxPageSize = 100
)

type TodoAPIController struct {
	app    *app.App
	logger *zerolog.Logger
}

func NewTodoAPIController(app *app.App, logger *zerolog.Logger) *TodoAPIController {
	return &TodoAPIController{app, logger}
}

func (c *TodoAPIController) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/todos", c.getTodos).Methods(http.MethodGet)
	r.HandleFunc("/todos", c.createTodo).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id:[0-9]+}/complete", c.completeTodo).Methods(http.MethodPost)
}

// Handlers

func (c *TodoAPIController) getTodos(w http.ResponseWriter, r *http.Request) {
	query, err := parseGetTodosQuery(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	todos, err := c.app.Queries.GetTodos.Handle(r.Context(), query)

	if err != nil {
		c.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, todos, query)
}

func (c *TodoAPIController) createTodo(w http.ResponseWriter, r *http.Request) {
	command, err := parseCreateTodoCommand(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := c.app.Commands.CreateTodo.Handle(r.Context(), command); err != nil {
		c.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteResponse(w, http.StatusAccepted, nil)
}

func (c *TodoAPIController) completeTodo(w http.ResponseWriter, r *http.Request) {
	command, err := parseCompleteTodoCommand(r)

	if err != nil {
		c.logError(err)
		goutils.WriteValidationError(w, err)
		return
	}

	if err := c.app.Commands.CompleteTodo.Handle(r.Context(), command); err != nil {
		c.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteResponse(w, http.StatusAccepted, nil)
}

// Input parsers

func parseGetTodosQuery(r *http.Request) (*query.GetTodos, error) {
	errorMap := map[string]string{}

	offset := 0
	limit := limitMaxPageSize

	var err error

	if value := r.FormValue(fieldPaginationOffset); value != "" {
		offset, err = strconv.Atoi(value)
		if err != nil || offset < 0 {
			errorMap[fieldPaginationOffset] = errCodeInvalidOffset
		}
	}

	if value := r.FormValue(fieldPaginationLimit); value != "" {
		limit, err = strconv.Atoi(value)
		if err != nil || limit < 1 || limitMaxPageSize < limit {
			errorMap[fieldPaginationLimit] = errCodeInvalidLimit
		}
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &query.GetTodos{
		Offset: offset,
		Limit:  limit,
	}, nil
}

func parseCreateTodoCommand(r *http.Request) (*command.CreateTodo, error) {
	errorMap := map[string]string{}

	todo := &domain.Todo{}
	if err := json.NewDecoder(r.Body).Decode(todo); err != nil {
		errorMap[fieldRequestBody] = errCodeInvalidRequestBody
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &command.CreateTodo{
		Todo: todo,
	}, nil
}

func parseCompleteTodoCommand(r *http.Request) (*command.CompleteTodo, error) {
	errorMap := map[string]string{}

	id, err := strconv.Atoi(mux.Vars(r)[fieldID])
	if err != nil {
		errorMap[fieldID] = errCodeInvalidID
	}

	if 0 < len(errorMap) {
		return nil, goutils.NewValidationError(errorMap)
	}

	return &command.CompleteTodo{
		ID: id,
	}, nil
}

// Utils

func (c *TodoAPIController) logError(err error) {
	c.logger.Error().Err(err).Send()
}
