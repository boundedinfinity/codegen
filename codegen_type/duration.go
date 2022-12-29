package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenTypeDuration struct {
	CodeGenTypeBase
	Min o.Option[CodeGenTypeDuration] `json:"min,omitempty"`
	Max o.Option[CodeGenTypeDuration] `json:"max,omitempty"`
}

func (t CodeGenTypeDuration) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined()
}

func (t CodeGenTypeDuration) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Duration
}

var _ CodeGenType = &CodeGenTypeDuration{}
