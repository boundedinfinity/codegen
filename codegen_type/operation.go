package codegen_type

import o "github.com/boundedinfinity/go-commoner/optioner"

type CodeGenProjectOperation struct {
	Root        o.Option[string] `json:"root,omitempty"`
	Path        o.Option[string] `json:"path,omitempty"`
	Name        o.Option[string] `json:"name,omitempty"`
	Description o.Option[string] `json:"description,omitempty"`
	Input       o.Option[string] `json:"input,omitempty"`
	Output      o.Option[string] `json:"output,omitempty"`
}
