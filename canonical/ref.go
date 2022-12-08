package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CanonicalRef struct {
	CanonicalBase
	Ref o.Option[string] `json:"ref,omitempty" yaml:"ref,omitempty"`
}

func (t CanonicalRef) CType() canonical_type.CanonicalType {
	return canonical_type.String
}

func (t CanonicalRef) HasValidation() bool {
	return true
}

func (t CanonicalRef) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Ref
}

var _ Canonical = &CanonicalRef{}
