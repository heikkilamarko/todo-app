package utils

import "github.com/xeipuuv/gojsonschema"

// SchemaValidator struct
type SchemaValidator struct {
	sl gojsonschema.JSONLoader
}

// NewSchemaValidator func
func NewSchemaValidator(schema string) *SchemaValidator {
	return &SchemaValidator{
		sl: gojsonschema.NewStringLoader(schema),
	}
}

// Validate method
func (v *SchemaValidator) Validate(data string) error {

	dl := gojsonschema.NewStringLoader(data)

	res, err := gojsonschema.Validate(v.sl, dl)
	if err != nil {
		return err
	}

	if !res.Valid() {
		m := map[string]string{}
		for _, e := range res.Errors() {
			m[e.Field()] = e.String()
		}
		return NewValidationError(m)
	}

	return nil
}
