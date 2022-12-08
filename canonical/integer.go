package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CanonicalInteger struct {
	CanonicalBase
	Minimum    o.Option[int] `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum    o.Option[int] `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	MultipleOf o.Option[int] `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
}

func (t CanonicalInteger) CType() canonical_type.CanonicalType {
	return canonical_type.Integer
}

func (t CanonicalInteger) HasValidation() bool {
	return t.Minimum.Defined() || t.Maximum.Defined() || t.MultipleOf.Defined()
}

func (t CanonicalInteger) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Integer
}

var _ Canonical = &CanonicalInteger{}
