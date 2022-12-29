package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenTypeEnum struct {
	CodeGenTypeBase
	Values        mapper.Mapper[string, string] `json:"values,omitempty"`
	CaseSensitive o.Option[bool]                `json:"caseSensitive,omitempty"`
}

func (t CodeGenTypeEnum) HasValidation() bool {
	return t.CaseSensitive.Defined()
}

func (t CodeGenTypeEnum) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Enum
}

var _ CodeGenType = &CodeGenTypeEnum{}
