package internal

type Config struct {
	App                string
	DBConnectionString string
	TemporalHostPort   string
	LogLevel           string
}

func (c *Config) Load() error {
	c.App = Env("APP_NAME", "")
	c.DBConnectionString = Env("APP_DB_CONNECTION_STRING", "")
	c.TemporalHostPort = Env("APP_TEMPORAL_HOSTPORT", "")
	c.LogLevel = Env("APP_LOG_LEVEL", "warn")

	return nil
}
