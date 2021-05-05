// Package utils provides utility functionality.
package utils

import (
	"github.com/xeipuuv/gojsonschema"
)

// JSONSchemaValidator struct
type JSONSchemaValidator struct {
	sl gojsonschema.JSONLoader
}

// NewJSONSchemaValidator func
func NewJSONSchemaValidator(s string) *JSONSchemaValidator {
	sl := gojsonschema.NewStringLoader(s)
	return &JSONSchemaValidator{sl}
}

func (v *JSONSchemaValidator) Validate(d string) ([]string, error) {
	dl := gojsonschema.NewStringLoader(d)

	result, err := gojsonschema.Validate(v.sl, dl)
	if err != nil {
		return nil, err
	}

	if !result.Valid() {
		var ves []string
		for _, ve := range result.Errors() {
			ves = append(ves, ve.String())
		}
		return ves, nil
	}

	return nil, nil
}
