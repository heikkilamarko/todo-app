package todos

import "github.com/nats-io/nats.go"

// Controller struct
type Controller struct {
	Repository Repository
	nc         *nats.EncodedConn
}

// NewController func
func NewController(r Repository, nc *nats.EncodedConn) *Controller {
	return &Controller{r, nc}
}
