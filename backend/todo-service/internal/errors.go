package internal

import "fmt"

type ValidationError struct {
	Errors map[string][]string
}

func (err ValidationError) Error() string {
	return fmt.Sprintf("%v", err.Errors)
}
