package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
)

// ([^@ \t\r\n]+)@([^@ \t\r\n]+\.[^@ \t\r\n]+)
// https://ihateregex.io/expr/email/

type CodeGenTypeEmail struct {
	SourceMeta
	RenderNamespace
	CodeGenTypeBase
}

func (t CodeGenTypeEmail) HasValidation() bool {
	return true
}

func (t CodeGenTypeEmail) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Email
}

func (t CodeGenTypeEmail) ValidateSchema() error {
	return nil
}

var _ CodeGenType = &CodeGenTypeEmail{}
