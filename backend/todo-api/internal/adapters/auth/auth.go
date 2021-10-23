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

func GetUserName(ctx context.Context) string {
	var c struct {
		Name string `mapstructure:"name"`
	}
	mapstructure.Decode(GetAccessToken(ctx), &c)
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
	mapstructure.Decode(GetAccessToken(ctx), &c)
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
