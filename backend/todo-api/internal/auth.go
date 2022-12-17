package internal

import "context"

type contextKey string

var ContextKeyAccessToken = contextKey("access_token")

func GetAccessToken(ctx context.Context) map[string]any {
	return ctx.Value(ContextKeyAccessToken).(map[string]any)
}
