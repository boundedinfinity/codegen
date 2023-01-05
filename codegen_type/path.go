package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
)

type CodeGenTypePath struct {
	SourceMeta
	RenderNamespace
	CodeGenTypeBase
}

func (t CodeGenTypePath) HasValidation() bool {
	return false
}

func (t CodeGenTypePath) SchemaType() codegen_type_id.CodgenTypeId {
	return codegen_type_id.Ref
}

func (t CodeGenTypePath) ValidateSchema() error {
	return nil
}

var _ CodeGenType = &CodeGenTypePath{}
