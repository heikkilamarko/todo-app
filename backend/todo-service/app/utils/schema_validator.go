package utils

import (
	"embed"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

// SchemaValidator struct
type SchemaValidator struct {
	fs      embed.FS
	schemas map[string]*gojsonschema.Schema
}

// NewSchemaValidator func
func NewSchemaValidator(fs embed.FS) *SchemaValidator {
	return &SchemaValidator{fs, make(map[string]*gojsonschema.Schema)}
}

// Validate method
func (v *SchemaValidator) Validate(schemaName string, doc []byte) error {

	schema, err := v.getSchema(schemaName)
	if err != nil {
		return err
	}

	r, err := schema.Validate(gojsonschema.NewStringLoader(string(doc)))
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

func (v *SchemaValidator) getSchema(schemaName string) (*gojsonschema.Schema, error) {

	schema, found := v.schemas[schemaName]

	if !found {
		schemaBytes, err := v.fs.ReadFile(fmt.Sprintf("schemas/%s.json", schemaName))
		if err != nil {
			return nil, ErrSchemaNotFound
		}

		sl := gojsonschema.NewSchemaLoader()
		sl.Draft = gojsonschema.Draft7
		sl.AutoDetect = false

		schema, err = sl.Compile(gojsonschema.NewStringLoader(string(schemaBytes)))
		if err != nil {
			return nil, ErrInvalidSchema
		}

		v.schemas[schemaName] = schema
	}

	return schema, nil
}
