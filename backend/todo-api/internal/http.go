package internal

import (
	"encoding/json"
	"errors"
	"net/http"
)

type DataResponse struct {
	Data any `json:"data"`
	Meta any `json:"meta,omitempty"`
}

type ErrorResponse struct {
	Error ErrorResponseError `json:"error"`
}

type ErrorResponseError struct {
	ErrorCode
	Details map[string][]string `json:"details,omitempty"`
}

func NewDataResponse(data, meta any) *DataResponse {
	return &DataResponse{data, meta}
}

func NewErrorResponse(code ErrorCode, details map[string][]string) *ErrorResponse {
	return &ErrorResponse{ErrorResponseError{code, details}}
}

func WriteErrorResponse(w http.ResponseWriter, code ErrorCode, err error) {
	var verr ValidationError
	if errors.As(err, &verr) {
		WriteResponse(w, http.StatusBadRequest, NewErrorResponse(code, verr.Errors))
	} else {
		WriteResponse(w, http.StatusInternalServerError, nil)
	}
}

func WriteResponse(w http.ResponseWriter, statusCode int, body any) {
	if body == nil {
		w.WriteHeader(statusCode)
		return
	}

	data, err := json.Marshal(body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}
