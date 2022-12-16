package codegen_project

import o "github.com/boundedinfinity/go-commoner/optioner"

type CodeGenProjectOperation struct {
	Description o.Option[string] `json:"description,omitempty" yaml:"description,omitempty"`
	Input       o.Option[string] `json:"input,omitempty" yaml:"input,omitempty"`
	Output      o.Option[string] `json:"output,omitempty" yaml:"output,omitempty"`
}
