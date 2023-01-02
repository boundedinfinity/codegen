package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	"github.com/boundedinfinity/go-jsonschema/model"
)

type CodeGenType interface {
	Source() *SourceMeta
	Base() *CodeGenTypeBase
	Namespace() *RenderNamespace
	SchemaType() codegen_type_id.CodgenTypeId
	HasValidation() bool
	ValidateSchema() error
}

type CodeGenTypeContext struct {
	FileInfo  SourceMeta
	Namespace RenderNamespace
	Schema    CodeGenType
}

func (t *CodeGenTypeContext) GetFileInfo() *SourceMeta {
	return &t.FileInfo
}

func (t *CodeGenTypeContext) GetNamespace() *RenderNamespace {
	return &t.Namespace
}

var _ LoaderContext = &CodeGenTypeContext{}

type JsonSchemaContext struct {
	FileInfo  SourceMeta
	Namespace RenderNamespace
	Schema    model.JsonSchema
}

func (t *JsonSchemaContext) GetFileInfo() *SourceMeta {
	return &t.FileInfo
}

func (t *JsonSchemaContext) GetNamespace() *RenderNamespace {
	return &t.Namespace
}

var _ LoaderContext = &JsonSchemaContext{}
