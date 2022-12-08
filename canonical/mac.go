package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"
)

// ([^@ \t\r\n]+)@([^@ \t\r\n]+\.[^@ \t\r\n]+)
// https://ihateregex.io/expr/email/

type CanonicalMac struct {
	CanonicalBase
}

func (t CanonicalMac) CType() canonical_type.CanonicalType {
	return canonical_type.Mac
}

func (t CanonicalMac) HasValidation() bool {
	return true
}

func (t CanonicalMac) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Mac
}

var _ Canonical = &CanonicalMac{}
