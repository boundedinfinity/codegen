package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"boundedinfinity/codegen/template_type"
	"text/template"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

type CodeGenProjectTemplates struct {
	Header o.Option[CodeGenProjectHeader] `json:"header,omitempty"`
	Files  []*CodeGenProjectTemplateFile  `json:"files,omitempty"`
}

type CodeGenProjectTemplateFile struct {
	Header  o.Option[CodeGenProjectHeader] `json:"header,omitempty"`
	Path    o.Option[string]               `json:"path,omitempty"`
	Content o.Option[string]               `json:"content,omitempty"`
}

type TemplateContext struct {
	FileInfo       LoaderFileInfo
	OutputMimeType mime_type.MimeType
	TemplateType   template_type.TemplateType
	TypeId         codegen_type_id.CodgenTypeId
	Template       *template.Template
}

func (t *TemplateContext) GetFileInfo() *LoaderFileInfo {
	return &t.FileInfo
}

var _ LoaderContext = &TemplateContext{}
