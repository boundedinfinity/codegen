package codegen_project

import "github.com/boundedinfinity/go-commoner/optioner/mapper"

type CodeGenProjectProject struct {
	Root       string
	Source     string
	Info       CodeGenProjectInfo                              `json:"info,omitempty" yaml:"info,omitempty"`
	Mappings   mapper.Mapper[string, string]                   `json:"mappings,omitempty" yaml:"mappings,omitempty"`
	Operations mapper.Mapper[string, *CodeGenProjectOperation] `json:"operations,omitempty" yaml:"operations,omitempty"`
	Templates  CodeGenProjectTemplates                         `json:"templates,omitempty" yaml:"templates,omitempty"`
	Schemas    []CodeGenProjectTypeFile                        `json:"schemas,omitempty" yaml:"schemas,omitempty"`
}

func NewProject() *CodeGenProjectProject {
	return &CodeGenProjectProject{
		Info:       NewInfo(),
		Mappings:   mapper.Mapper[string, string]{},
		Operations: mapper.Mapper[string, *CodeGenProjectOperation]{},
	}
}
