package render_context

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

type RenderContext interface {
	Base() *RenderContextBase
	HasValidation() bool
}

type RenderContextBase struct {
	Root        string
	Source      string
	Id          string
	RootNs      string
	CurrNs      string
	SchemaNs    string
	RelNs       string
	SchemaType  codegen_type_id.CodgenTypeId
	MimeType    mime_type.MimeType
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
