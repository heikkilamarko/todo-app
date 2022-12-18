package internal

type Config struct {
	App                          string
	Address                      string
	DBConnectionString           string
	NATSURL                      string
	NATSToken                    string
	CentrifugoTokenHMACSecretKey []byte
	LogLevel                     string
	AuthIssuer                   string
	AuthClaimIss                 string
	AuthClaimAud                 string
	AuthZBackend                 string
}

func (c *Config) Load() error {
	c.App = Env("APP_NAME", "")
	c.Address = Env("APP_ADDRESS", ":8080")
	c.DBConnectionString = Env("APP_DB_CONNECTION_STRING", "")
	c.NATSURL = Env("APP_NATS_URL", "")
	c.NATSToken = Env("APP_NATS_TOKEN", "")
	c.CentrifugoTokenHMACSecretKey = EnvBytes("APP_CENTRIFUGO_TOKEN_HMAC_SECRET_KEY", nil)
	c.LogLevel = Env("APP_LOG_LEVEL", "warn")
	c.AuthIssuer = Env("APP_AUTH_ISSUER", "")
	c.AuthClaimIss = Env("APP_AUTH_CLAIM_ISS", "")
	c.AuthClaimAud = Env("APP_AUTH_CLAIM_AUD", "")
	c.AuthZBackend = Env("APP_AUTHZ_BACKEND", "opa")

	return nil
}
