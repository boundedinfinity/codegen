package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
)

// ^[\+]?[(]?([0-9]{3})[)]?[-\s\.]?([0-9]{3})[-\s\.]?([0-9]{4,6})$
// https://ihateregex.io/expr/e164-phone
// https://ihateregex.io/expr/phone

type CodeGenTypePhone struct {
	CodeGenTypeBase
}

func (t CodeGenTypePhone) HasValidation() bool {
	return true
}

func (t CodeGenTypePhone) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Phone
}

var _ CodeGenType = &CodeGenTypePhone{}
