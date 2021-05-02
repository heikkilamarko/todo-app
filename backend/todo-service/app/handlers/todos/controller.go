package todos

import (
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog"
)

// Controller struct
type Controller struct {
	Repository Repository
	NatsConn   *nats.Conn
	Logger     *zerolog.Logger
}

// NewController func
func NewController(r Repository, c *nats.Conn, l *zerolog.Logger) *Controller {
	return &Controller{r, c, l}
}
