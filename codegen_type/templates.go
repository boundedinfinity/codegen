package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"boundedinfinity/codegen/codegen_type/template_type"
	"text/template"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

type CodeGenProjectTemplates struct {
	Header o.Option[CodeGenTemplateHeader] `json:"header,omitempty"`
	Files  []*CodeGenProjectTemplateFile   `json:"files,omitempty"`
}

type CodeGenProjectTemplateFile struct {
	SourceMeta
	RenderNamespace
	Header  o.Option[CodeGenTemplateHeader] `json:"header,omitempty"`
	Content o.Option[string]                `json:"content,omitempty"`
}

type TemplateMeta struct {
	SourceMeta
	RenderNamespace
	OutputMimeType mime_type.MimeType
	TemplateType   template_type.TemplateType
	TypeId         codegen_type_id.CodgenTypeId
	Template       *template.Template
}

var _ LoaderContext = &CodeGenProjectTemplateFile{}
var _ LoaderContext = &TemplateMeta{}
