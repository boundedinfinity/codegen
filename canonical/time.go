package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

// (?:(?:31(\/|-|\.)(?:0?[13578]|1[02]))\1|(?:(?:29|30)(\/|-|\.)(?:0?[13-9]|1[0-2])\2))(?:(?:1[6-9]|[2-9]\d)?\d{2})$|^(?:29(\/|-|\.)0?2\3(?:(?:(?:1[6-9]|[2-9]\d)?(?:0[48]|[2468][048]|[13579][26])|(?:(?:16|[2468][048]|[3579][26])00))))$|^(?:0?[1-9]|1\d|2[0-8])(\/|-|\.)(?:(?:0?[1-9])|(?:1[0-2]))\4(?:(?:1[6-9]|[2-9]\d)?\d{2})
// https://ihateregex.io/expr/date

type CanonicalTime struct {
	CanonicalBase
	Before o.Option[CanonicalDate]     `json:"before,omitempty" yaml:"before,omitempty"`
	After  o.Option[CanonicalDate]     `json:"after,omitempty" yaml:"after,omitempty"`
	Within o.Option[CanonicalDuration] `json:"within,omitempty" yaml:"within,omitempty"`
	Ahead  o.Option[CanonicalDuration] `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Behind o.Option[CanonicalDuration] `json:"maximum,omitempty" yaml:"maximum,omitempty"`
}

func (t CanonicalTime) CType() canonical_type.CanonicalType {
	return canonical_type.Time
}

func (t CanonicalTime) HasValidation() bool {
	return t.Before.Defined() || t.After.Defined() || t.Ahead.Defined() ||
		t.Behind.Defined() || t.Within.Defined()
}

func (t CanonicalTime) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Time
}

var _ Canonical = &CanonicalTime{}
