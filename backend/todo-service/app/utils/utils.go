// Package utils provides utility functionality.
package utils

import "os"

// Env func
func Env(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
