package auth

import (
	"context"

	"github.com/mitchellh/mapstructure"
)

type contextKey string

var ContextKeyAccessToken = contextKey("access_token")

func GetAccessToken(ctx context.Context) map[string]interface{} {
	return ctx.Value(ContextKeyAccessToken).(map[string]interface{})
}

func GetSubject(ctx context.Context) string {
	var c struct {
		Subject string `mapstructure:"sub"`
	}
	_ = mapstructure.Decode(GetAccessToken(ctx), &c)
	return c.Subject
}

func GetUserName(ctx context.Context) string {
	var c struct {
		Name string `mapstructure:"name"`
	}
	_ = mapstructure.Decode(GetAccessToken(ctx), &c)
	return c.Name
}

func GetRoles(ctx context.Context) []string {
	var c struct {
		ResourceAccess struct {
			TodoAPI struct {
				Roles []string `mapstructure:"roles"`
			} `mapstructure:"todo-api"`
		} `mapstructure:"resource_access"`
	}
	_ = mapstructure.Decode(GetAccessToken(ctx), &c)
	return c.ResourceAccess.TodoAPI.Roles
}

func IsInRole(ctx context.Context, role string) bool {
	for _, r := range GetRoles(ctx) {
		if r == role {
			return true
		}
	}
	return false
}

func IsInAnyRole(ctx context.Context, roles ...string) bool {
	for _, r := range GetRoles(ctx) {
		for _, role := range roles {
			if r == role {
				return true
			}
		}
	}
	return false
}
