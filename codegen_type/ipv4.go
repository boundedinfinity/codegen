package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

// https://ihateregex.io/expr/ip/

type CodeGenTypeIpv4 struct {
	CodeGenTypeBase
	Minimum o.Option[string] `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum o.Option[string] `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Mask    o.Option[string] `json:"mask,omitempty" yaml:"mask,omitempty"`
}

func (t CodeGenTypeIpv4) HasValidation() bool {
	return t.Minimum.Defined() || t.Maximum.Defined() || t.Mask.Defined()
}

func (t CodeGenTypeIpv4) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Ipv4
}

var _ CodeGenType = &CodeGenTypeIpv4{}