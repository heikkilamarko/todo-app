package internal

import "context"

type contextKey string

var (
	ContextKeyAccessToken = contextKey("access_token")
	ContextKeyAuthZResult = contextKey("authz_result")
)

func GetAccessToken(ctx context.Context) map[string]any {
	return ctx.Value(ContextKeyAccessToken).(map[string]any)
}

func GetAuthZResult(ctx context.Context) *AuthZResult {
	return ctx.Value(ContextKeyAuthZResult).(*AuthZResult)
}
