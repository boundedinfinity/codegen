package conical

import (
	"boundedinfinity/codegen/conical/conical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type ConicalRef struct {
	ConicalBase
	Minimum o.Option[int] `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum o.Option[int] `json:"maximum,omitempty" yaml:"maximum,omitempty"`
}

func (t ConicalRef) CType() conical_type.ConicalType {
	return conical_type.String
}

func (t ConicalRef) HasValidation() bool {
	return true
}

var _ Conical = &ConicalRef{}
