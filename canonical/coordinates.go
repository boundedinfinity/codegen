package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"
)

// https://ihateregex.io/expr/lat-long

type CanonicalCoordinate struct {
	CanonicalBase
}

func (t CanonicalCoordinate) CType() canonical_type.CanonicalType {
	return canonical_type.Coordinate
}

func (t CanonicalCoordinate) HasValidation() bool {
	return true
}

func (t CanonicalCoordinate) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Coordinate
}

var _ Canonical = &CanonicalCoordinate{}
