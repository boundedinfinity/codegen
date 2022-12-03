package conical

import (
	"boundedinfinity/codegen/conical/conical_type"
)

type ConicalObject struct {
	ConicalBase
	Properties []Conical `json:"properties,omitempty" yaml:"properties,omitempty"`
}

func (t ConicalObject) CType() conical_type.ConicalType {
	return conical_type.Object
}

func (t ConicalObject) HasValidation() bool {
	for _, property := range t.Properties {
		if property.HasValidation() {
			return true
		}
	}

	return false
}

var _ Conical = &ConicalObject{}
