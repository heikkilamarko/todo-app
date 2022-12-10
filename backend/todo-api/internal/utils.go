package internal

import (
	"net/http"
	"os"
	"strings"
)

func Env(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}

func EnvBytes(key string, fallback []byte) []byte {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return []byte(value)
}

func TokenFromHeader(r *http.Request) string {
	a := r.Header.Get("Authorization")
	if 7 < len(a) && strings.ToUpper(a[0:6]) == "BEARER" {
		return a[7:]
	}
	return ""
}
