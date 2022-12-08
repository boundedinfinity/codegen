package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CanonicalDateTime struct {
	CanonicalBase
	Before o.Option[CanonicalDate]     `json:"before,omitempty" yaml:"before,omitempty"`
	After  o.Option[CanonicalDate]     `json:"after,omitempty" yaml:"after,omitempty"`
	Within o.Option[CanonicalDuration] `json:"within,omitempty" yaml:"within,omitempty"`
	Ahead  o.Option[CanonicalDuration] `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Behind o.Option[CanonicalDuration] `json:"maximum,omitempty" yaml:"maximum,omitempty"`
}

func (t CanonicalDateTime) CType() canonical_type.CanonicalType {
	return canonical_type.Date
}

func (t CanonicalDateTime) HasValidation() bool {
	return t.Before.Defined() || t.After.Defined() || t.Ahead.Defined() ||
		t.Behind.Defined() || t.Within.Defined()
}

func (t CanonicalDateTime) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Datetime
}

var _ Canonical = &CanonicalDateTime{}
