package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CanonicalDuration struct {
	CanonicalBase
	Min o.Option[CanonicalDuration] `json:"min,omitempty" yaml:"min,omitempty"`
	Max o.Option[CanonicalDuration] `json:"max,omitempty" yaml:"max,omitempty"`
}

func (t CanonicalDuration) CType() canonical_type.CanonicalType {
	return canonical_type.Duration
}

func (t CanonicalDuration) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined()
}

func (t CanonicalDuration) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Duration
}

var _ Canonical = &CanonicalDuration{}
