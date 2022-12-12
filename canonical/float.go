package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CanonicalFloat struct {
	CanonicalBase
	Min        o.Option[float64] `json:"min,omitempty" yaml:"min,omitempty"`
	Max        o.Option[float64] `json:"max,omitempty" yaml:"max,omitempty"`
	MultipleOf o.Option[float64] `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
}

func (t CanonicalFloat) CType() canonical_type.CanonicalType {
	return canonical_type.Float
}

func (t CanonicalFloat) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined() || t.MultipleOf.Defined()
}

func (t CanonicalFloat) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Float
}

var _ Canonical = &CanonicalFloat{}
