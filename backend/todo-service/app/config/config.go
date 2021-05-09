// Package config provides application configuration functionality.
package config

import (
	"encoding/json"
	"os"

	"github.com/rs/zerolog"
)

// Config struct
type Config struct {
	App                string        `json:"app"`
	DBConnectionString string        `json:"db_connection_string"`
	NATSUrl            string        `json:"nats_url"`
	NATSToken          string        `json:"nats_token"`
	LogLevel           zerolog.Level `json:"log_level"`
}

// New func
func New() *Config {
	return &Config{}
}

const redactedString = "[redacted]"

// String method
func (c *Config) String() string {
	cc := *c
	cc.DBConnectionString = redactedString
	cc.NATSToken = redactedString
	if b, err := json.Marshal(cc); err == nil {
		return string(b)
	}
	return ""
}

// Load method
func (c *Config) Load() {
	c.App = getEnv("APP_NAME", "")
	c.DBConnectionString = getEnv("APP_DB_CONNECTION_STRING", "")
	c.NATSUrl = getEnv("APP_NATS_URL", "")
	c.NATSToken = getEnv("APP_NATS_TOKEN", "")

	var err error

	if c.LogLevel, err = zerolog.ParseLevel(getEnv("APP_LOG_LEVEL", "warn")); err != nil {
		c.LogLevel = zerolog.WarnLevel
	}
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
