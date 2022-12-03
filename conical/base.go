package conical

import (
	"boundedinfinity/codegen/conical/conical_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type Conical interface {
	CType() conical_type.ConicalType
	HasValidation() bool
}

type ConicalBase struct {
	Import      string           `json:"import,omitempty" yaml:"import,omitempty"`
	Type        string           `json:"type,omitempty" yaml:"type,omitempty"`
	Package     string           `json:"package,omitempty" yaml:"package,omitempty"`
	Name        string           `json:"name,omitempty" yaml:"name,omitempty"`
	Source      string           `json:"source,omitempty" yaml:"source,omitempty"`
	Description o.Option[string] `json:"description,omitempty" yaml:"description,omitempty"`
	Pointer     o.Option[bool]   `json:"pointer,omitempty" yaml:"pointer,omitempty"`
	Array       o.Option[bool]   `json:"array,omitempty" yaml:"array,omitempty"`
	Exported    o.Option[bool]   `json:"exported,omitempty" yaml:"exported,omitempty"`
}
