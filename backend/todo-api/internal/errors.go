package internal

import (
	"errors"
	"fmt"
)

var ErrUnauthorized = errors.New("unauthorized")
var ErrNotFound = errors.New("not found")

type ValidationError struct {
	Errors map[string][]string
}

func (err ValidationError) Error() string {
	return fmt.Sprintf("%v", err.Errors)
}
