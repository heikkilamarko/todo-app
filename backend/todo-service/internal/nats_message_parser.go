package internal

import (
	"encoding/json"

	"github.com/nats-io/nats.go/jetstream"
)

type NATSMessageParser struct {
	Validator *SchemaValidator
}

func (p *NATSMessageParser) Parse(msg jetstream.Msg, model any) error {
	if err := p.Validator.Validate(msg.Subject(), msg.Data()); err != nil {
		return err
	}

	if err := json.Unmarshal(msg.Data(), model); err != nil {
		return err
	}

	return nil
}
