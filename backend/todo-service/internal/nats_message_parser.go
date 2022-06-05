package internal

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type NATSMessageParser struct {
	Validator *SchemaValidator
}

func (p *NATSMessageParser) Parse(m *nats.Msg, model any) error {
	if err := p.Validator.Validate(m.Subject, m.Data); err != nil {
		return err
	}

	if err := json.Unmarshal(m.Data, model); err != nil {
		return err
	}

	return nil
}
