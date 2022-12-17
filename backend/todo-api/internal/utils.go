package internal

import "os"

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
