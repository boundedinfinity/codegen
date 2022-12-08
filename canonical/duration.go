package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CanonicalDuration struct {
	CanonicalBase
	Minimum o.Option[CanonicalDuration] `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum o.Option[CanonicalDuration] `json:"maximum,omitempty" yaml:"maximum,omitempty"`
}

func (t CanonicalDuration) CType() canonical_type.CanonicalType {
	return canonical_type.Duration
}

func (t CanonicalDuration) HasValidation() bool {
	return t.Minimum.Defined() || t.Maximum.Defined()
}

func (t CanonicalDuration) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Duration
}

var _ Canonical = &CanonicalDuration{}
