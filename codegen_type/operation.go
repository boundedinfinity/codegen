package codegen_type

import o "github.com/boundedinfinity/go-commoner/optioner"

type CodeGenProjectOperation struct {
	Root        o.Option[string] `json:"root,omitempty" yaml:"root,omitempty"`
	Path        o.Option[string] `json:"path,omitempty" yaml:"path,omitempty"`
	Name        o.Option[string] `json:"name,omitempty" yaml:"name,omitempty"`
	Description o.Option[string] `json:"description,omitempty" yaml:"description,omitempty"`
	Input       o.Option[string] `json:"input,omitempty" yaml:"input,omitempty"`
	Output      o.Option[string] `json:"output,omitempty" yaml:"output,omitempty"`
}
