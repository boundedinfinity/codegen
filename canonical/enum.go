package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CanonicalEnum struct {
	CanonicalBase
	Values        mapper.Mapper[string, string] `json:"values,omitempty" yaml:"values,omitempty"`
	CaseSensitive o.Option[bool]                `json:"caseSensitive,omitempty" yaml:"caseSensitive,omitempty"`
}

func (t CanonicalEnum) CType() canonical_type.CanonicalType {
	return canonical_type.String
}

func (t CanonicalEnum) HasValidation() bool {
	return t.CaseSensitive.Defined()
}

func (t CanonicalEnum) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Enum
}

var _ Canonical = &CanonicalEnum{}
