package codegen_type

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenProjectProject struct {
	Info       CodeGenProjectInfo            `json:"info,omitempty"`
	Mappings   mapper.Mapper[string, string] `json:"mappings,omitempty"`
	Operations []*CodeGenProjectOperation    `json:"operations,omitempty"`
	Templates  CodeGenProjectTemplates       `json:"templates,omitempty"`
	Schemas    []*CodeGenProjectTypeFile     `json:"schemas,omitempty"`
}

func NewProject() *CodeGenProjectProject {
	return &CodeGenProjectProject{
		Info:       NewInfo(),
		Mappings:   mapper.Mapper[string, string]{},
		Operations: make([]*CodeGenProjectOperation, 0),
	}
}

type ProjectContext struct {
	FileInfo LoaderFileInfo
	Project  CodeGenProjectProject
}

func (t *ProjectContext) GetFileInfo() *LoaderFileInfo {
	return &t.FileInfo
}

var _ LoaderContext = &ProjectContext{}

type OperationContext struct {
	*ProjectContext
	Name        o.Option[string] `json:"name,omitempty"`
	Description o.Option[string] `json:"description,omitempty"`
	Input       o.Option[string] `json:"input,omitempty"`
	Output      o.Option[string] `json:"output,omitempty"`
}

func (t *OperationContext) GetFileInfo() *LoaderFileInfo {
	return &t.FileInfo
}

var _ LoaderContext = &OperationContext{}
