package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenTypeString struct {
	CodeGenTypeBase
	Min   o.Option[int]    `json:"min,omitempty" yaml:"min,omitempty"`
	Max   o.Option[int]    `json:"max,omitempty" yaml:"max,omitempty"`
	Regex o.Option[string] `json:"regex,omitempty" yaml:"regex,omitempty"`
}

func (t CodeGenTypeString) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined() || t.Regex.Defined()
}

func (t CodeGenTypeString) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.String
}

var _ CodeGenType = &CodeGenTypeString{}
