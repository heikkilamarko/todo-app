// Package schemas provides schema validation functionality.
package schemas

import (
	"github.com/xeipuuv/gojsonschema"
)

type Validator struct {
	sl gojsonschema.JSONLoader
}

func NewValidator(s string) *Validator {
	sl := gojsonschema.NewStringLoader(s)
	return &Validator{sl}
}

func (v *Validator) Validate(d string) ([]string, error) {
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
