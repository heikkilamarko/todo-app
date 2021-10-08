package auth

import "context"

type contextKey string

var ContextKeyAccessToken = contextKey("access_token")

func GetAccessToken(ctx context.Context) map[string]interface{} {
	return ctx.Value(ContextKeyAccessToken).(map[string]interface{})
}

func GetUserName(token map[string]interface{}) string {
	switch v := token["name"].(type) {
	case string:
		return v
	default:
		return ""
	}
}

func GetRoles(token map[string]interface{}) []string {
	var roles []string

	switch v1 := token["resource_access"].(type) {
	case map[string]interface{}:
		switch v2 := v1["todo-api"].(type) {
		case map[string]interface{}:
			switch v3 := v2["roles"].(type) {
			case []interface{}:
				for _, role := range v3 {
					roles = append(roles, role.(string))
				}
			}
		}
	}

	return roles
}
