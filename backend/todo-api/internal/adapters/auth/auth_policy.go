package auth

import (
	"errors"
	"net/http"
)

const (
	roleUser   = "todo-user"
	roleViewer = "todo-viewer"
)

var ErrUnauthorized = errors.New("unauthorized")

func AuthorizeWrite(r *http.Request) error {
	if IsInRole(r.Context(), roleUser) {
		return nil
	}
	return ErrUnauthorized
}

func AuthorizeRead(r *http.Request) error {
	if IsInAnyRole(r.Context(), roleUser, roleViewer) {
		return nil
	}
	return ErrUnauthorized
}
