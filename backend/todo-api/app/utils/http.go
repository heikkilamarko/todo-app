package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

// DataResponse struct
type DataResponse struct {
	Meta interface{} `json:"meta,omitempty"`
	Data interface{} `json:"data"`
}

// NewDataResponse func
func NewDataResponse(data, meta interface{}) *DataResponse {
	return &DataResponse{meta, data}
}

// ErrorResponse struct
type ErrorResponse struct {
	Error *ErrorResponseError `json:"error"`
}

// ErrorResponseError struct
type ErrorResponseError struct {
	Code    string            `json:"code"`
	Details map[string]string `json:"details,omitempty"`
}

// NewErrorResponse func
func NewErrorResponse(code string, details map[string]string) *ErrorResponse {
	return &ErrorResponse{
		Error: &ErrorResponseError{code, details},
	}
}

// NewBadRequestResponse func
func NewBadRequestResponse(details map[string]string) *ErrorResponse {
	return NewErrorResponse(ErrCodeBadRequest, details)
}

// NewUnauthorizedResponse func
func NewUnauthorizedResponse(details map[string]string) *ErrorResponse {
	return NewErrorResponse(ErrCodeUnauthorized, details)
}

// NewNotFoundResponse func
func NewNotFoundResponse(details map[string]string) *ErrorResponse {
	return NewErrorResponse(ErrCodeNotFound, details)
}

// NewInternalErrorResponse func
func NewInternalErrorResponse(details map[string]string) *ErrorResponse {
	return NewErrorResponse(ErrCodeInternalError, details)
}

// WriteOK writes 200 response
func WriteOK(w http.ResponseWriter, data, meta interface{}) {
	WriteResponse(w, http.StatusOK, NewDataResponse(data, meta))
}

// WriteCreated writes 201 response
func WriteCreated(w http.ResponseWriter, data, meta interface{}) {
	WriteResponse(w, http.StatusCreated, NewDataResponse(data, meta))
}

// WriteNoContent writes 204 response
func WriteNoContent(w http.ResponseWriter) {
	WriteResponse(w, http.StatusNoContent, nil)
}

// WriteBadRequest writes 400 response
func WriteBadRequest(w http.ResponseWriter, details map[string]string) {
	WriteResponse(w, http.StatusBadRequest, NewBadRequestResponse(details))
}

// WriteUnauthorized writes 401 response
func WriteUnauthorized(w http.ResponseWriter, details map[string]string) {
	WriteResponse(w, http.StatusUnauthorized, NewUnauthorizedResponse(details))
}

// WriteNotFound writes 404 response
func WriteNotFound(w http.ResponseWriter, details map[string]string) {
	WriteResponse(w, http.StatusNotFound, NewNotFoundResponse(details))
}

// WriteInternalError writes 500 response
func WriteInternalError(w http.ResponseWriter, details map[string]string) {
	WriteResponse(w, http.StatusInternalServerError, NewInternalErrorResponse(details))
}

// WriteValidationError writes 400 or 500 response
func WriteValidationError(w http.ResponseWriter, err error) {
	var vErr *ValidationError
	if errors.As(err, &vErr) {
		WriteBadRequest(w, vErr.ErrorMap)
	} else {
		WriteInternalError(w, nil)
	}
}

// WriteResponse func
func WriteResponse(w http.ResponseWriter, code int, body interface{}) {
	if body != nil {
		content, err := json.Marshal(body)

		if err != nil {
			WriteInternalError(w, nil)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(content)
	} else {
		w.WriteHeader(code)
	}
}
