package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CanonicalFloat struct {
	CanonicalBase
	Minimum    o.Option[int] `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum    o.Option[int] `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	MultipleOf o.Option[int] `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
}

func (t CanonicalFloat) CType() canonical_type.CanonicalType {
	return canonical_type.Float
}

func (t CanonicalFloat) HasValidation() bool {
	return t.Minimum.Defined() || t.Maximum.Defined() || t.MultipleOf.Defined()
}

func (t CanonicalFloat) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Float
}

var _ Canonical = &CanonicalFloat{}
