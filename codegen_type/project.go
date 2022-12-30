package codegen_type

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenProject struct {
	Info       CodeGenInfo                   `json:"info,omitempty"`
	Mappings   mapper.Mapper[string, string] `json:"mappings,omitempty"`
	Operations []*CodeGenProjectOperation    `json:"operations,omitempty"`
	Templates  CodeGenProjectTemplates       `json:"templates,omitempty"`
	Schemas    []*CodeGenProjectTypeFile     `json:"schemas,omitempty"`
}

func NewProject() *CodeGenProject {
	return &CodeGenProject{
		Info:       NewInfo(),
		Mappings:   mapper.Mapper[string, string]{},
		Operations: make([]*CodeGenProjectOperation, 0),
	}
}

type ProjectContext struct {
	FileInfo  FileInfo
	Namespace Namespace
	Project   CodeGenProject
}

func (t *ProjectContext) GetFileInfo() *FileInfo {
	return &t.FileInfo
}

func (t *ProjectContext) GetNamespace() *Namespace {
	return &t.Namespace
}

var _ LoaderContext = &ProjectContext{}

type OperationContext struct {
	*ProjectContext
	Name        o.Option[string] `json:"name,omitempty"`
	Description o.Option[string] `json:"description,omitempty"`
	Input       o.Option[string] `json:"input,omitempty"`
	Output      o.Option[string] `json:"output,omitempty"`
}

func (t *OperationContext) GetFileInfo() *FileInfo {
	return &t.FileInfo
}

var _ LoaderContext = &OperationContext{}
