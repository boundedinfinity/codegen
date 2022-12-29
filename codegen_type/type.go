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
	FileInfo LoaderFileInfo
	Schema   CodeGenType
}

func (t *CodeGenTypeContext) GetFileInfo() *LoaderFileInfo {
	return &t.FileInfo
}

var _ LoaderContext = &CodeGenTypeContext{}

type JsonSchemaContext struct {
	FileInfo LoaderFileInfo
	Schema   model.JsonSchema
}

func (t *JsonSchemaContext) GetFileInfo() *LoaderFileInfo {
	return &t.FileInfo
}

var _ LoaderContext = &JsonSchemaContext{}
