package internal

import "net/http"

const (
	roleUser   = "todo-user"
	roleViewer = "todo-viewer"
)

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
