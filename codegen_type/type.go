package codegen_type

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
)

type CodeGenProjectTypeFile struct {
	Path o.Option[string] `json:"path,omitempty"`
	Root o.Option[string] `json:"root,omitempty"`
}

type CodeGenTypeContext struct {
	FileInfo  FileInfo
	Namespace Namespace
	Schema    CodeGenType
}

func (t *CodeGenTypeContext) GetFileInfo() *FileInfo {
	return &t.FileInfo
}

func (t *CodeGenTypeContext) GetNamespace() *Namespace {
	return &t.Namespace
}

var _ LoaderContext = &CodeGenTypeContext{}

type JsonSchemaContext struct {
	FileInfo  FileInfo
	Namespace Namespace
	Schema    model.JsonSchema
}

func (t *JsonSchemaContext) GetFileInfo() *FileInfo {
	return &t.FileInfo
}

func (t *JsonSchemaContext) GetNamespace() *Namespace {
	return &t.Namespace
}

var _ LoaderContext = &JsonSchemaContext{}
