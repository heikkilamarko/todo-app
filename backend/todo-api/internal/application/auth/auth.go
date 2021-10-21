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

func GetUserName(token map[string]interface{}) string {
	var c struct {
		Name string `mapstructure:"name"`
	}
	mapstructure.Decode(token, &c)
	return c.Name
}

func GetRoles(token map[string]interface{}) []string {
	var c struct {
		ResourceAccess struct {
			TodoAPI struct {
				Roles []string `mapstructure:"roles"`
			} `mapstructure:"todo-api"`
		} `mapstructure:"resource_access"`
	}
	mapstructure.Decode(token, &c)
	return c.ResourceAccess.TodoAPI.Roles
}

func IsInRole(token map[string]interface{}, role string) bool {
	for _, r := range GetRoles(token) {
		if r == role {
			return true
		}
	}
	return false
}
