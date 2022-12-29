package codegen_type

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
)

type CodeGenProjectTypeFile struct {
	Path o.Option[string] `json:"path,omitempty" yaml:"path,omitempty"`
	Root o.Option[string] `json:"root,omitempty" yaml:"root,omitempty"`
}

var _ LoaderContext = &TypeLoaderContext{}

type TypeLoaderContext struct {
	FileInfo LoaderFileInfo
	Schema   CodeGenType
}

func (t *TypeLoaderContext) GetFileInfo() *LoaderFileInfo {
	return &t.FileInfo
}

type JsonSchemaLoaderContext struct {
	FileInfo LoaderFileInfo
	Schema   model.JsonSchema
}

func (t *JsonSchemaLoaderContext) GetFileInfo() *LoaderFileInfo {
	return &t.FileInfo
}

var _ LoaderContext = &JsonSchemaLoaderContext{}
