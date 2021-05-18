package utils

import "errors"

var (
	// ErrSchemaNotFound error
	ErrSchemaNotFound = errors.New("schema not found")
	// ErrInvalidSchema error
	ErrInvalidSchema = errors.New("invalid schema")
)
