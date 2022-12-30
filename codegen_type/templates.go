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
	Header  o.Option[CodeGenTemplateHeader] `json:"header,omitempty"`
	Path    o.Option[string]                `json:"path,omitempty"`
	Content o.Option[string]                `json:"content,omitempty"`
}

type TemplateContext struct {
	FileInfo       FileInfo
	Namespace      Namespace
	OutputMimeType mime_type.MimeType
	TemplateType   template_type.TemplateType
	TypeId         codegen_type_id.CodgenTypeId
	Template       *template.Template
}

func (t *TemplateContext) GetFileInfo() *FileInfo {
	return &t.FileInfo
}

func (t *TemplateContext) GetNamespace() *Namespace {
	return &t.Namespace
}

var _ LoaderContext = &TemplateContext{}
