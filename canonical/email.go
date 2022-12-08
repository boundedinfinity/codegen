package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"
)

// ([^@ \t\r\n]+)@([^@ \t\r\n]+\.[^@ \t\r\n]+)
// https://ihateregex.io/expr/email/

type CanonicalEmail struct {
	CanonicalBase
}

func (t CanonicalEmail) CType() canonical_type.CanonicalType {
	return canonical_type.Email
}

func (t CanonicalEmail) HasValidation() bool {
	return true
}

func (t CanonicalEmail) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Email
}

var _ Canonical = &CanonicalEmail{}
