package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	"github.com/boundedinfinity/go-jsonschema/model"
)

type CodeGenType interface {
	Source() *SourceMeta
	Namespace() *RenderNamespace
	SchemaType() codegen_type_id.CodgenTypeId
	Base() *CodeGenTypeBase
	HasValidation() bool
	ValidateSchema() error
}

type JsonSchemaContext struct {
	SourceMeta
	RenderNamespace
	Schema model.JsonSchema
}

var _ LoaderContext = &JsonSchemaContext{}
