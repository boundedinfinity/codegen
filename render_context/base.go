package render_context

import "boundedinfinity/codegen/codegen_type/codegen_type_id"

type RenderContext interface {
	Base() *RenderContextBase
}

type RenderContextBase struct {
	OutputPath    string
	Root          string
	Source        string
	Id            string
	RootNs        string
	CurrNs        string
	SchemaNs      string
	RelNs         string
	SchemaType    codegen_type_id.CodgenTypeId
	Name          string
	Description   string
	IsPublic      bool
	IsRequired    bool
	IsInterface   bool
	Header        string
	HasValidation bool
}

func (t *RenderContextBase) Base() *RenderContextBase {
	return t
}
