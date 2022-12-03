package conical

import (
	"boundedinfinity/codegen/conical/conical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type ConicalDuration struct {
	ConicalBase
	Minimum o.Option[ConicalDuration] `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum o.Option[ConicalDuration] `json:"maximum,omitempty" yaml:"maximum,omitempty"`
}

func (t ConicalDuration) CType() conical_type.ConicalType {
	return conical_type.Duration
}

func (t ConicalDuration) HasValidation() bool {
	return t.Minimum.Defined() || t.Maximum.Defined()
}

var _ Conical = &ConicalDuration{}
