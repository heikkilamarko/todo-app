package utils

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

// MessageParser struct
type MessageParser struct {
	v *SchemaValidator
}

// NewMessageParser func
func NewMessageParser(v *SchemaValidator) *MessageParser {
	return &MessageParser{v}
}

// Parse method
func (p *MessageParser) Parse(message *nats.Msg, model interface{}) error {

	if err := p.v.Validate(message.Subject, message.Data); err != nil {
		return err
	}

	if err := json.Unmarshal(message.Data, model); err != nil {
		return err
	}

	return nil
}
