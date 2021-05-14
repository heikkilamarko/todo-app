package utils

import (
	"encoding/json"
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
