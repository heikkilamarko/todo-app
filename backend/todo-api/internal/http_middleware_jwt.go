package internal

import (
	"context"
	"crypto/tls"
	"log/slog"
	"net/http"
	"strings"

	"github.com/hashicorp/cap/jwt"
)

type JWTMiddlewareConfig struct {
	Issuer     string
	Audience   []string
	ContextKey any
	Logger     *slog.Logger
}

func JWTMiddleware(ctx context.Context, config *JWTMiddlewareConfig) func(next http.Handler) http.Handler {
	http.DefaultClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	keySet, err := jwt.NewOIDCDiscoveryKeySet(ctx, config.Issuer, "")
	if err != nil {
		panic(err)
	}

	validator, err := jwt.NewValidator(keySet)
	if err != nil {
		panic(err)
	}

	expected := jwt.Expected{
		Issuer:    config.Issuer,
		Audiences: config.Audience,
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := tokenFromHeader(r)
			if token == "" {
				config.Logger.Error("token is empty")
				WriteResponse(w, http.StatusUnauthorized, nil)
				return
			}

			claims, err := validator.Validate(r.Context(), token, expected)
			if err != nil {
				config.Logger.Error(err.Error())
				WriteResponse(w, http.StatusUnauthorized, nil)
				return
			}

			r = r.WithContext(context.WithValue(r.Context(), config.ContextKey, claims))

			next.ServeHTTP(w, r)
		})
	}
}

func tokenFromHeader(r *http.Request) string {
	a := r.Header.Get("Authorization")
	if 7 < len(a) && strings.ToUpper(a[0:6]) == "BEARER" {
		return a[7:]
	}
	return ""
}
