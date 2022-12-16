package codegen_project

import (
	"boundedinfinity/codegen/header_strategy"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenProjectHeader struct {
	Content  o.Option[string]                         `json:"content,omitempty" yaml:"content,omitempty"`
	Path     o.Option[string]                         `json:"path,omitempty" yaml:"path,omitempty"`
	Strategy o.Option[header_strategy.HeaderStrategy] `json:"strategy,omitempty" yaml:"strategy,omitempty"`
}
