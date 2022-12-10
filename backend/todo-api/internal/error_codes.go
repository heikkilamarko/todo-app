package internal

type ErrorCode struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var ErrCodeInvalidRequest = ErrorCode{
	Code:    "400",
	Message: "invalid request",
}
