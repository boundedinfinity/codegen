package codegen_type

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenTypeBase struct {
	SourceMeta
	RenderNamespace
	Id          o.Option[string] `json:"id,omitempty"`
	Name        o.Option[string] `json:"name,omitempty"`
	Description o.Option[string] `json:"description,omitempty"`
	Public      o.Option[bool]   `json:"public,omitempty"`
	Required    o.Option[bool]   `json:"required,omitempty"`
	Deprecated  o.Option[bool]   `json:"deprecated,omitempty"`
}

func (t *CodeGenTypeBase) Base() *CodeGenTypeBase {
	return t
}

func (t CodeGenTypeBase) SchemaId() o.Option[string] {
	return t.Id
}

var _ LoaderContext = &CodeGenTypeBase{}
