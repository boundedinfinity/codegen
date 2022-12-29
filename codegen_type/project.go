package codegen_type

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenProjectProject struct {
	Info       CodeGenProjectInfo            `json:"info,omitempty" yaml:"info,omitempty"`
	Mappings   mapper.Mapper[string, string] `json:"mappings,omitempty" yaml:"mappings,omitempty"`
	Operations []*CodeGenProjectOperation    `json:"operations,omitempty" yaml:"operations,omitempty"`
	Templates  CodeGenProjectTemplates       `json:"templates,omitempty" yaml:"templates,omitempty"`
	Schemas    []*CodeGenProjectTypeFile     `json:"schemas,omitempty" yaml:"schemas,omitempty"`
}

func NewProject() *CodeGenProjectProject {
	return &CodeGenProjectProject{
		Info:       NewInfo(),
		Mappings:   mapper.Mapper[string, string]{},
		Operations: make([]*CodeGenProjectOperation, 0),
	}
}

type ProjectLoaderContext struct {
	FileInfo LoaderFileInfo
	Project  CodeGenProjectProject
}

func (t *ProjectLoaderContext) GetFileInfo() *LoaderFileInfo {
	return &t.FileInfo
}

var _ LoaderContext = &ProjectLoaderContext{}

type OperationLoaderContext struct {
	*ProjectLoaderContext
	Name        o.Option[string] `json:"name,omitempty" yaml:"name,omitempty"`
	Description o.Option[string] `json:"description,omitempty" yaml:"description,omitempty"`
	Input       o.Option[string] `json:"input,omitempty" yaml:"input,omitempty"`
	Output      o.Option[string] `json:"output,omitempty" yaml:"output,omitempty"`
}

func (t *OperationLoaderContext) GetFileInfo() *LoaderFileInfo {
	return &t.FileInfo
}

var _ LoaderContext = &OperationLoaderContext{}
