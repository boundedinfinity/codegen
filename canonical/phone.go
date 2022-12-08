package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"
)

// ^[\+]?[(]?([0-9]{3})[)]?[-\s\.]?([0-9]{3})[-\s\.]?([0-9]{4,6})$
// https://ihateregex.io/expr/e164-phone
// https://ihateregex.io/expr/phone

type CanonicalPhone struct {
	CanonicalBase
}

func (t CanonicalPhone) CType() canonical_type.CanonicalType {
	return canonical_type.Phone
}

func (t CanonicalPhone) HasValidation() bool {
	return true
}

func (t CanonicalPhone) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Phone
}

var _ Canonical = &CanonicalPhone{}
