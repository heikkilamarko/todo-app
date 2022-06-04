package internal

type ErrorCode struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var ErrCodeInvalidRequest = ErrorCode{
	Code:    "E1001",
	Message: "invalid request",
}
