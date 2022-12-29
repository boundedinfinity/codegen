package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenTypeInteger struct {
	CodeGenTypeBase
	Min        o.Option[int64] `json:"min,omitempty"`
	Max        o.Option[int64] `json:"max,omitempty"`
	MultipleOf o.Option[int64] `json:"multipleOf,omitempty"`
}

func (t CodeGenTypeInteger) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined() || t.MultipleOf.Defined()
}

func (t CodeGenTypeInteger) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Integer
}

var _ CodeGenType = &CodeGenTypeInteger{}
