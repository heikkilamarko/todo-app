package adapters

import (
	"context"
	"encoding/json"
	"todo-service/internal/domain"

	"github.com/centrifugal/gocent/v3"
)

const (
	subjectTodoCreateOk      = "todo.create.ok"
	subjectTodoCreateError   = "todo.create.error"
	subjectTodoCompleteOk    = "todo.complete.ok"
	subjectTodoCompleteError = "todo.complete.error"
)

type todoCreateOkMessage struct {
	Todo *domain.Todo `json:"todo"`
}

type todoCompleteOkMessage struct {
	ID int `json:"id"`
}

type errorMessage struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}

type messageWrapper struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type TodoCentrifugoMessagePublisher struct {
	c *gocent.Client
}

func NewTodoCentrifugoMessagePublisher(addr, key string) *TodoCentrifugoMessagePublisher {
	return &TodoCentrifugoMessagePublisher{
		gocent.New(gocent.Config{Addr: addr, Key: key}),
	}
}

func (mp *TodoCentrifugoMessagePublisher) TodoCreateOk(ctx context.Context, todo *domain.Todo) error {
	return mp.publish(
		ctx,
		subjectTodoCreateOk,
		&todoCreateOkMessage{
			todo,
		},
	)
}

func (mp *TodoCentrifugoMessagePublisher) TodoCreateError(ctx context.Context, message string) error {
	return mp.publish(
		ctx,
		subjectTodoCreateError,
		&errorMessage{
			Code:    subjectTodoCreateError,
			Message: message,
		},
	)
}

func (mp *TodoCentrifugoMessagePublisher) TodoCompleteOk(ctx context.Context, id int) error {
	return mp.publish(
		ctx,
		subjectTodoCompleteOk,
		&todoCompleteOkMessage{
			id,
		},
	)
}

func (mp *TodoCentrifugoMessagePublisher) TodoCompleteError(ctx context.Context, message string) error {
	return mp.publish(
		ctx,
		subjectTodoCompleteError,
		&errorMessage{
			Code:    subjectTodoCompleteError,
			Message: message,
		},
	)
}

func (mp *TodoCentrifugoMessagePublisher) publish(ctx context.Context, messageType string, message interface{}) error {
	data, err := json.Marshal(messageWrapper{messageType, message})
	if err != nil {
		return err
	}

	_, err = mp.c.Publish(ctx, "notifications", data)
	return err
}
