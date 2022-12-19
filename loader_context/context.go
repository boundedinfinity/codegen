package loader_context

import (
	"boundedinfinity/codegen/codegen_project"
	"boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"boundedinfinity/codegen/template_type"
	"text/template"

	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

type LoaderContext interface {
	GetFileInfo() *LoaderFileInfo
}

type LoaderFileInfo struct {
	Source   string
	Root     string
	IsFile   bool
	MimeType mime_type.MimeType
}

type ProjectLoaderContext struct {
	FileInfo LoaderFileInfo
	Project  codegen_project.CodeGenProjectProject
}

func (t *ProjectLoaderContext) GetFileInfo() *LoaderFileInfo {
	return &t.FileInfo
}

var _ LoaderContext = &ProjectLoaderContext{}

type TypeLoaderContext struct {
	FileInfo LoaderFileInfo
	Schema   codegen_type.CodeGenType
}

func (t *TypeLoaderContext) GetFileInfo() *LoaderFileInfo {
	return &t.FileInfo
}

var _ LoaderContext = &TypeLoaderContext{}

type JsonSchemaLoaderContext struct {
	FileInfo LoaderFileInfo
	Schema   model.JsonSchema
}

func (t *JsonSchemaLoaderContext) GetFileInfo() *LoaderFileInfo {
	return &t.FileInfo
}

var _ LoaderContext = &JsonSchemaLoaderContext{}

type TemplateLoaderContext struct {
	FileInfo       LoaderFileInfo
	OutputMimeType mime_type.MimeType
	TemplateType   template_type.TemplateType
	TypeId         codegen_type_id.CodgenTypeId
	Template       *template.Template
}

func (t *TemplateLoaderContext) GetFileInfo() *LoaderFileInfo {
	return &t.FileInfo
}

var _ LoaderContext = &TemplateLoaderContext{}
