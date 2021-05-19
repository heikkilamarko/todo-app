// Package config provides application configuration functionality.
package config

import "todo-api/app/utils"

// Config struct
type Config struct {
	App                string `json:"app"`
	Address            string `json:"address"`
	DBConnectionString string `json:"db_connection_string"`
	NATSUrl            string `json:"nats_url"`
	NATSToken          string `json:"nats_token"`
	LogLevel           string `json:"log_level"`
}

// Load func
func Load() *Config {
	return &Config{
		App:                utils.Env("APP_NAME", ""),
		Address:            utils.Env("APP_ADDRESS", ":8080"),
		DBConnectionString: utils.Env("APP_DB_CONNECTION_STRING", ""),
		NATSUrl:            utils.Env("APP_NATS_URL", ""),
		NATSToken:          utils.Env("APP_NATS_TOKEN", ""),
		LogLevel:           utils.Env("APP_LOG_LEVEL", "warn"),
	}
}
