package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CanonicalString struct {
	CanonicalBase
	Min o.Option[int] `json:"min,omitempty" yaml:"min,omitempty"`
	Max o.Option[int] `json:"max,omitempty" yaml:"max,omitempty"`
}

func (t CanonicalString) CType() canonical_type.CanonicalType {
	return canonical_type.String
}

func (t CanonicalString) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined()
}

var _ Canonical = &CanonicalString{}
