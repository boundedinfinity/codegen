package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CanonicalInteger struct {
	CanonicalBase
	Min        o.Option[int64] `json:"min,omitempty" yaml:"min,omitempty"`
	Max        o.Option[int64] `json:"max,omitempty" yaml:"max,omitempty"`
	MultipleOf o.Option[int64] `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
}

func (t CanonicalInteger) CType() canonical_type.CanonicalType {
	return canonical_type.Integer
}

func (t CanonicalInteger) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined() || t.MultipleOf.Defined()
}

func (t CanonicalInteger) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Integer
}

var _ Canonical = &CanonicalInteger{}
