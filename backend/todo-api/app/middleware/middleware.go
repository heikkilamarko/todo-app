// Package middleware provides middleware functionality.
package middleware

import (
	"context"
	"net/http"
	"time"
	"todo-api/app/utils"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

// NotFoundHandler func
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.WriteNotFound(w, nil)
}

// Logger middleware
func Logger(logger *zerolog.Logger) func(next http.Handler) http.Handler {
	return hlog.NewHandler(*logger)
}

// RequestLogger middleware
func RequestLogger() func(next http.Handler) http.Handler {
	return hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Send()
	})
}

// ErrorRecovery middleware
func ErrorRecovery() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					hlog.FromRequest(r).Error().Msgf("%s", err)
					utils.WriteInternalError(w, nil)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

// APIKey middleware
func APIKey(apiKey, headerKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if ak := r.Header.Get(headerKey); ak != apiKey {
				utils.WriteUnauthorized(w, nil)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// Timeout middleware
func Timeout(duration time.Duration) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), duration)
			defer cancel()
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
