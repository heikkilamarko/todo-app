package internal

import "net/http"

func NotFound(w http.ResponseWriter, r *http.Request) {
	WriteResponse(w, http.StatusNotFound, nil)
}
