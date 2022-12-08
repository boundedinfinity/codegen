package canonical

import (
	"boundedinfinity/codegen/canonical/canonical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

// https://ihateregex.io/expr/ip/

type CanonicalIpv6 struct {
	CanonicalBase
	Minimum o.Option[string] `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum o.Option[string] `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Mask    o.Option[string] `json:"mask,omitempty" yaml:"mask,omitempty"`
}

func (t CanonicalIpv6) CType() canonical_type.CanonicalType {
	return canonical_type.Ipv6
}

func (t CanonicalIpv6) HasValidation() bool {
	return t.Minimum.Defined() || t.Maximum.Defined() || t.Mask.Defined()
}

func (t CanonicalIpv6) SchemaType() canonical_type.CanonicalType {
	return canonical_type.Ipv6
}

var _ Canonical = &CanonicalIpv6{}
