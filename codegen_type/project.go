package codegen_type

import (
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenProject struct {
	SourceMeta
	RenderNamespace
	Info       CodeGenInfo                   `json:"info,omitempty"`
	Mappings   mapper.Mapper[string, string] `json:"mappings,omitempty"`
	Operations []*CodeGenProjectOperation    `json:"operations,omitempty"`
	Templates  CodeGenProjectTemplates       `json:"templates,omitempty"`
	Types      []CodeGenType                 `json:"types,omitempty"`
}

func NewProject() *CodeGenProject {
	return &CodeGenProject{
		Info:       NewInfo(),
		Mappings:   mapper.Mapper[string, string]{},
		Operations: make([]*CodeGenProjectOperation, 0),
	}
}

var _ LoaderContext = &CodeGenProject{}
