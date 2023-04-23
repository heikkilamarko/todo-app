package internal

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/exp/slog"
)

type AuthZMiddlewareConfig struct {
	AuthZ      AuthZ
	Permission string
	ContextKey any
	Logger     *slog.Logger
}

func AuthZMiddleware(ctx context.Context, config *AuthZMiddlewareConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ar, err := config.AuthZ.Authorize(r.Context(), &AuthZQuery{
				Token:      GetAccessToken(r.Context()),
				Permission: config.Permission,
			})

			if err != nil {
				config.Logger.Error(err.Error())
				WriteResponse(w, http.StatusForbidden, nil)
				return
			}

			config.Logger.Info(fmt.Sprintf("[%s][%s] %#v", r.URL.Path, config.Permission, ar))

			if !ar.Allow {
				WriteResponse(w, http.StatusForbidden, nil)
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), config.ContextKey, ar))

			next.ServeHTTP(w, r)
		})
	}
}
