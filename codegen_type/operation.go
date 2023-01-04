package codegen_type

import o "github.com/boundedinfinity/go-commoner/optioner"

type CodeGenProjectOperation struct {
	SourceMeta
	RenderNamespace
	Name        o.Option[string] `json:"name,omitempty"`
	Description o.Option[string] `json:"description,omitempty"`
	Input       o.Option[string] `json:"input,omitempty"`
	Output      o.Option[string] `json:"output,omitempty"`
}

var _ LoaderContext = &CodeGenProjectOperation{}
