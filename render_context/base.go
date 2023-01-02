package render_context

import (
	"boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
)

type RenderContext interface {
	Base() *RenderContextBase
	HasValidation() bool
}

type RenderContextBase struct {
	codegen_type.RenderNamespace
	codegen_type.SourceMeta
	Id          string
	SchemaType  codegen_type_id.CodgenTypeId
	Name        string
	Description string
	IsPublic    bool
	IsRequired  bool
	IsInterface bool
	Header      string
}

func (t *RenderContextBase) Base() *RenderContextBase {
	return t
}
