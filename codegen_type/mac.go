package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
)

// ([^@ \t\r\n]+)@([^@ \t\r\n]+\.[^@ \t\r\n]+)
// https://ihateregex.io/expr/email/

type CodeGenTypeMac struct {
	CodeGenTypeBase
}

func (t CodeGenTypeMac) HasValidation() bool {
	return true
}

func (t CodeGenTypeMac) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Mac
}

func (t CodeGenTypeMac) ValidateSchema() error {
	return nil
}

var _ CodeGenType = &CodeGenTypeMac{}
