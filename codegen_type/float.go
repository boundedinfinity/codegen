package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenTypeFloat struct {
	CodeGenTypeBase
	Min        o.Option[float64] `json:"min,omitempty"`
	Max        o.Option[float64] `json:"max,omitempty"`
	MultipleOf o.Option[float64] `json:"multipleOf,omitempty"`
}

func (t CodeGenTypeFloat) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined() || t.MultipleOf.Defined()
}

func (t CodeGenTypeFloat) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Float
}

var _ CodeGenType = &CodeGenTypeFloat{}
