package conical

import (
	"boundedinfinity/codegen/conical/conical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type ConicalDate struct {
	ConicalBase
	Before   o.Option[ConicalDate]     `json:"before,omitempty" yaml:"before,omitempty"`
	After    o.Option[ConicalDate]     `json:"after,omitempty" yaml:"after,omitempty"`
	Duration o.Option[ConicalDuration] `json:"duration,omitempty" yaml:"duration,omitempty"`
	Ahead    o.Option[ConicalDuration] `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Behind   o.Option[ConicalDuration] `json:"maximum,omitempty" yaml:"maximum,omitempty"`
}

func (t ConicalDate) CType() conical_type.ConicalType {
	return conical_type.Date
}

func (t ConicalDate) HasValidation() bool {
	return t.Before.Defined() || t.After.Defined() || t.Ahead.Defined() || t.Behind.Defined()
}

var _ Conical = &ConicalDate{}
