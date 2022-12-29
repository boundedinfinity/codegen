package codegen_type

import (
	"boundedinfinity/codegen/header_strategy"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenProjectHeader struct {
	Content  o.Option[string]                         `json:"content,omitempty"`
	Path     o.Option[string]                         `json:"path,omitempty"`
	Strategy o.Option[header_strategy.HeaderStrategy] `json:"strategy,omitempty"`
}
