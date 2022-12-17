package internal

import "context"

type AuthZQuery struct {
	Permission string `json:"permission"`
	Token      any    `json:"token"`
}

type AuthZResult struct {
	Allow       bool
	Sub         string
	Username    string
	Permissions []string
}

type AuthZ interface {
	Authorize(ctx context.Context, query *AuthZQuery) (*AuthZResult, error)
}
