package conical

import (
	"boundedinfinity/codegen/conical/conical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type ConicalString struct {
	ConicalBase
	Minimum o.Option[int] `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum o.Option[int] `json:"maximum,omitempty" yaml:"maximum,omitempty"`
}

func (t ConicalString) CType() conical_type.ConicalType {
	return conical_type.String
}

func (t ConicalString) HasValidation() bool {
	return t.Minimum.Defined() || t.Maximum.Defined()
}

var _ Conical = &ConicalString{}
