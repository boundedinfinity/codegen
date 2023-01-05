package codegen_type

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenProject struct {
	SourceMeta
	RenderNamespace
	Info       CodeGenInfo                   `json:"info,omitempty"`
	Mappings   mapper.Mapper[string, string] `json:"mappings,omitempty"`
	Operations []CodeGenProjectOperation     `json:"operations,omitempty"`
	Templates  CodeGenProjectTemplates       `json:"templates,omitempty"`
	Types      []CodeGenType                 `json:"types,omitempty"`
}

func NewProject() *CodeGenProject {
	return &CodeGenProject{
		Info:       NewInfo(),
		Mappings:   mapper.Mapper[string, string]{},
		Operations: make([]CodeGenProjectOperation, 0),
	}
}

var _ LoaderContext = &CodeGenProject{}

type codeGenProject struct {
	SourceMeta
	RenderNamespace
	Info       CodeGenInfo                   `json:"info,omitempty"`
	Mappings   mapper.Mapper[string, string] `json:"mappings,omitempty"`
	Operations []CodeGenProjectOperation     `json:"operations,omitempty"`
	Templates  CodeGenProjectTemplates       `json:"templates,omitempty"`
	Types      []json.RawMessage             `json:"types,omitempty"`
}

func (t *CodeGenProject) UnmarshalJSON(data []byte) error {
	var d codeGenProject

	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	t.SourceMeta = d.SourceMeta
	t.RenderNamespace = d.RenderNamespace
	t.Info = d.Info
	t.Operations = d.Operations
	t.Templates = d.Templates

	for _, raw := range d.Types {
		var typ CodeGenType

		if err := UnmarshalJson(raw, &typ); err != nil {
			return err
		}

		t.Types = append(t.Types, typ)
	}

	return nil
}
