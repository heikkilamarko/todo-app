// Package config provides application configuration functionality.
package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Config struct
type Config struct {
	App                string        `json:"app"`
	Address            string        `json:"address"`
	DBConnectionString string        `json:"db_connection_string"`
	NATSUrl            string        `json:"nats_url"`
	NATSToken          string        `json:"nats_token"`
	LogLevel           zerolog.Level `json:"log_level"`
	RequestTimeout     time.Duration `json:"request_timeout"`
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
	c.Address = getEnv("APP_ADDRESS", ":8080")
	c.DBConnectionString = getEnv("APP_DB_CONNECTION_STRING", "")
	c.NATSUrl = getEnv("APP_NATS_URL", "")
	c.NATSToken = getEnv("APP_NATS_TOKEN", "")

	var err error

	if c.LogLevel, err = zerolog.ParseLevel(getEnv("APP_LOG_LEVEL", "warn")); err != nil {
		c.LogLevel = zerolog.WarnLevel
	}

	if c.RequestTimeout, err = time.ParseDuration(getEnv("APP_REQUEST_TIMEOUT", "10s")); err != nil {
		c.RequestTimeout = 10 * time.Second
	}
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
