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
	Type             o.Option[codegen_type_id.CodgenTypeId]
	OutputMimeType   mime_type.MimeType
	OutputExt        string
	TemplateType     template_type.TemplateType
	TemplateMimeTime mime_type.MimeType
	TemplateExt      string
	Template         *template.Template
}

type CodeGenProjectTemplateFile struct {
	TemplateMeta
	Header  o.Option[CodeGenTemplateHeader] `json:"header,omitempty"`
	Content o.Option[string]                `json:"content,omitempty"`
}

type CodeGenProjectTemplates struct {
	Header     o.Option[CodeGenTemplateHeader] `json:"header,omitempty"`
	Types      []*CodeGenProjectTemplateFile   `json:"types,omitempty"`
	TypeList   []*CodeGenProjectTemplateFile   `json:"type-list,omitempty"`
	Operations []*CodeGenProjectTemplateFile   `json:"operations,omitempty"`
}

var _ LoaderContext = &CodeGenProjectTemplateFile{}
var _ LoaderContext = &TemplateMeta{}
