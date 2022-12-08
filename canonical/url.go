package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"
)

// // https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()!@:%_\+.~#?&\/\/=]*)
// https://ihateregex.io/expr/url

type CanonicalUrl struct {
	CanonicalBase
}

func (t CanonicalUrl) CType() canonical_type.CanonicalType {
	return canonical_type.Url
}

func (t CanonicalUrl) HasValidation() bool {
	return true
}

func (t CanonicalUrl) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Url
}

var _ Canonical = &CanonicalUrl{}
