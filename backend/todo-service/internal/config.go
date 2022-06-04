package internal

type Config struct {
	App                string
	DBConnectionString string
	NATSURL            string
	NATSToken          string
	CentrifugoURL      string
	CentrifugoKey      string
	LogLevel           string
}

func (c *Config) Load() error {
	c.App = Env("APP_NAME", "")
	c.DBConnectionString = Env("APP_DB_CONNECTION_STRING", "")
	c.NATSURL = Env("APP_NATS_URL", "")
	c.NATSToken = Env("APP_NATS_TOKEN", "")
	c.CentrifugoURL = Env("APP_CENTRIFUGO_URL", "")
	c.CentrifugoKey = Env("APP_CENTRIFUGO_KEY", "")
	c.LogLevel = Env("APP_LOG_LEVEL", "warn")

	return nil
}
