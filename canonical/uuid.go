package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"
	"boundedinfinity/codegen/canonical/uuid_version"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

// https://ihateregex.io/expr/uuid

type CanonicalUuid struct {
	CanonicalBase
	CaseSensitive o.Option[bool]                     `json:"caseSensitive,omitempty" yaml:"caseSensitive,omitempty"`
	Version       o.Option[uuid_version.UuidVersion] `json:"version,omitempty" yaml:"version,omitempty"`
}

func (t CanonicalUuid) CType() canonical_type.CanonicalType {
	return canonical_type.Uuid
}

func (t CanonicalUuid) HasValidation() bool {
	return t.CaseSensitive.Defined()
}

func (t CanonicalUuid) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Uuid
}

var _ Canonical = &CanonicalUuid{}
