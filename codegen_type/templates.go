package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"boundedinfinity/codegen/codegen_type/template_type"
	"text/template"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

type TemplateMeta struct {
	SourceMeta
	RenderNamespace
	Type           codegen_type_id.CodgenTypeId
	OutputMimeType mime_type.MimeType
	TemplateType   template_type.TemplateType
	Template       *template.Template
}

type CodeGenProjectTemplateFile struct {
	SourceMeta
	RenderNamespace
	TemplateMeta
	Header  o.Option[CodeGenTemplateHeader] `json:"header,omitempty"`
	Content o.Option[string]                `json:"content,omitempty"`
}

type CodeGenProjectOperationTemplateFile struct {
	SourceMeta
	TemplateMeta
	RenderNamespace
	Header o.Option[CodeGenTemplateHeader] `json:"header,omitempty"`
}

type CodeGenProjectTemplates struct {
	Header     o.Option[CodeGenTemplateHeader] `json:"header,omitempty"`
	Types      []*CodeGenProjectTemplateFile   `json:"types,omitempty"`
	TypeList   []*CodeGenProjectTemplateFile   `json:"type-list,omitempty"`
	Operations []*CodeGenProjectTemplateFile   `json:"operations,omitempty"`
}

var _ LoaderContext = &CodeGenProjectTemplateFile{}
var _ LoaderContext = &TemplateMeta{}
