package conical

import (
	"boundedinfinity/codegen/conical/conical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type ConicalInteger struct {
	ConicalBase
	Minimum    o.Option[int] `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum    o.Option[int] `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	MultipleOf o.Option[int] `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
}

func (t ConicalInteger) CType() conical_type.ConicalType {
	return conical_type.Integer
}

func (t ConicalInteger) HasValidation() bool {
	return t.Minimum.Defined() || t.Maximum.Defined() || t.MultipleOf.Defined()
}

var _ Conical = &ConicalObject{}
