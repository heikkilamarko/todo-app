// Package constants provides application constants.
package constants

const (
	// ErrCodeInvalidPayload constant
	ErrCodeInvalidPayload = "invalid_payload"

	// ErrCodeInvalidTodoID constant
	ErrCodeInvalidTodoID = "invalid_todo_id"

	// ErrCodeInvalidOffset constant
	ErrCodeInvalidOffset = "invalid_offset"

	// ErrCodeInvalidLimit constant
	ErrCodeInvalidLimit = "invalid_limit"
)

const (
	// FieldID constant
	FieldID = "id"

	// FieldPaginationOffset constant
	FieldPaginationOffset = "offset"

	// FieldPaginationLimit constant
	FieldPaginationLimit = "limit"

	// FieldRequestBody constant
	FieldRequestBody = "request_body"
)

const (
	// PaginationLimitMax constant
	PaginationLimitMax = 100
)

const (
	// MessageTodoCreated constant
	MessageTodoCreated = "todo.created"

	// MessageTodoCompleted constant
	MessageTodoCompleted = "todo.completed"
)
