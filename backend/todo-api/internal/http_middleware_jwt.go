package internal

import (
	"context"
	"errors"
	"net/http"

	"github.com/hashicorp/cap/jwt"
	"github.com/rs/zerolog"
)

type JWTMiddlewareConfig struct {
	Issuer     string
	Iss        string
	Aud        []string
	ContextKey any
	Logger     *zerolog.Logger
}

func JWTMiddleware(ctx context.Context, config *JWTMiddlewareConfig) func(next http.Handler) http.Handler {
	keySet, err := jwt.NewOIDCDiscoveryKeySet(ctx, config.Issuer, "")
	if err != nil {
		panic(err)
	}

	validator, err := jwt.NewValidator(keySet)
	if err != nil {
		panic(err)
	}

	expected := jwt.Expected{
		Issuer:    config.Iss,
		Audiences: config.Aud,
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := TokenFromHeader(r)
			if token == "" {
				config.Logger.Error().Err(errors.New("token is empty")).Send()
				WriteResponse(w, http.StatusUnauthorized, nil)
				return
			}

			claims, err := validator.Validate(r.Context(), token, expected)
			if err != nil {
				config.Logger.Error().Err(err).Send()
				WriteResponse(w, http.StatusUnauthorized, nil)
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), config.ContextKey, claims))

			next.ServeHTTP(w, r)
		})
	}
}
