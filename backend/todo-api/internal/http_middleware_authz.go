package internal

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"
)

type AuthZMiddlewareConfig struct {
	AuthZ      AuthZ
	Permission string
	ContextKey any
	Logger     *zerolog.Logger
}

func AuthZMiddleware(ctx context.Context, config *AuthZMiddlewareConfig) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ar, err := config.AuthZ.Authorize(r.Context(), &AuthZQuery{
				Token:      GetAccessToken(r.Context()),
				Permission: config.Permission,
			})

			if err != nil {
				config.Logger.Error().Err(err).Send()
				WriteResponse(w, http.StatusForbidden, nil)
				return
			}

			config.Logger.Info().Msgf("[%s][%s] %#v", r.URL.Path, config.Permission, ar)

			if !ar.Allow {
				WriteResponse(w, http.StatusForbidden, nil)
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), config.ContextKey, ar))

			next.ServeHTTP(w, r)
		})
	}
}
