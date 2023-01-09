package codegen_type

import (
	type_visibility "boundedinfinity/codegen/codegen_type/type_visability"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenTypeBase struct {
	SourceMeta
	RenderNamespace
	Id          o.Option[string]                         `json:"id,omitempty"`
	Name        o.Option[string]                         `json:"name,omitempty"`
	Description o.Option[string]                         `json:"description,omitempty"`
	Visibility  o.Option[type_visibility.TypeVisibility] `json:"visibility,omitempty"`
	Required    o.Option[bool]                           `json:"required,omitempty"`
	Deprecated  o.Option[bool]                           `json:"deprecated,omitempty"`
}

func (t *CodeGenTypeBase) Base() *CodeGenTypeBase {
	return t
}

func (t CodeGenTypeBase) SchemaId() o.Option[string] {
	return t.Id
}

func (t CodeGenTypeBase) HasValidation() bool {
	return false
}

var _ LoaderContext = &CodeGenTypeBase{}
