package utils

import (
	"encoding/json"
	"errors"
)

var (
	// ErrNotFound error
	ErrNotFound = errors.New("not found")
	// ErrInternalError error
	ErrInternalError = errors.New("internal error")
)

// ValidationError error
type ValidationError struct {
	ErrorMap map[string]string
}

// NewValidationError func
func NewValidationError(errorMap map[string]string) *ValidationError {
	return &ValidationError{errorMap}
}

// Error method
func (v *ValidationError) Error() string {
	message, err := json.Marshal(v.ErrorMap)
	if err != nil {
		return ""
	}
	return string(message)
}

// ErrorMessage struct
type ErrorMessage struct {
	Code    string `json:"code"`
	Message string `json:"message,omitempty"`
}
