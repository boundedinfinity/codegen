package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenTypeRef struct {
	CodeGenTypeBase
	Ref o.Option[string] `json:"ref,omitempty" yaml:"ref,omitempty"`
}

func (t CodeGenTypeRef) HasValidation() bool {
	return true
}

func (t CodeGenTypeRef) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Ref
}

var _ CodeGenType = &CodeGenTypeRef{}