package utils

import (
	"encoding/json"
)

// MessageParser struct
type MessageParser struct {
	validator *SchemaValidator
}

// NewMessageParser func
func NewMessageParser(v *SchemaValidator) *MessageParser {
	return &MessageParser{v}
}

// Parse method
func (p *MessageParser) Parse(message []byte, model interface{}) error {

	if p.validator != nil {
		if err := p.validator.Validate(string(message)); err != nil {
			return err
		}
	}

	if err := json.Unmarshal(message, model); err != nil {
		return err
	}

	return nil
}
