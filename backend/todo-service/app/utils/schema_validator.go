package utils

import "github.com/xeipuuv/gojsonschema"

// SchemaValidator struct
type SchemaValidator struct {
	schema *gojsonschema.Schema
}

// Validate method
func (v *SchemaValidator) Validate(doc string) error {

	r, err := v.schema.Validate(gojsonschema.NewStringLoader(doc))
	if err != nil {
		return err
	}

	if !r.Valid() {
		m := map[string]string{}
		for _, e := range r.Errors() {
			m[e.Field()] = e.Description()
		}
		return NewValidationError(m)
	}

	return nil
}

// NewSchemaValidator func
func NewSchemaValidator(schema string) (*SchemaValidator, error) {

	sl := gojsonschema.NewSchemaLoader()
	sl.Draft = gojsonschema.Draft7
	sl.AutoDetect = false

	s, err := sl.Compile(gojsonschema.NewStringLoader(schema))

	if err != nil {
		return nil, err
	}

	return &SchemaValidator{s}, nil
}
