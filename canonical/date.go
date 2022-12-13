package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

// (?:(?:31(\/|-|\.)(?:0?[13578]|1[02]))\1|(?:(?:29|30)(\/|-|\.)(?:0?[13-9]|1[0-2])\2))(?:(?:1[6-9]|[2-9]\d)?\d{2})$|^(?:29(\/|-|\.)0?2\3(?:(?:(?:1[6-9]|[2-9]\d)?(?:0[48]|[2468][048]|[13579][26])|(?:(?:16|[2468][048]|[3579][26])00))))$|^(?:0?[1-9]|1\d|2[0-8])(\/|-|\.)(?:(?:0?[1-9])|(?:1[0-2]))\4(?:(?:1[6-9]|[2-9]\d)?\d{2})
// https://ihateregex.io/expr/date

type CanonicalDate struct {
	CanonicalBase
	Before o.Option[CanonicalDate]     `json:"before,omitempty" yaml:"before,omitempty"`
	After  o.Option[CanonicalDate]     `json:"after,omitempty" yaml:"after,omitempty"`
	Within o.Option[CanonicalDuration] `json:"within,omitempty" yaml:"within,omitempty"`
	Ahead  o.Option[CanonicalDuration] `json:"ahead,omitempty" yaml:"ahead,omitempty"`
	Behind o.Option[CanonicalDuration] `json:"behind,omitempty" yaml:"behind,omitempty"`
}

func (t CanonicalDate) CType() canonical_type.CanonicalType {
	return canonical_type.Date
}

func (t CanonicalDate) HasValidation() bool {
	return t.Before.Defined() || t.After.Defined() || t.Ahead.Defined() ||
		t.Behind.Defined() || t.Within.Defined()
}

func (t CanonicalDate) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Date
}

var _ Canonical = &CanonicalDate{}
