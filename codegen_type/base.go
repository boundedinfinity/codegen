package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenType interface {
	HasValidation() bool
	SchemaType() codegen_type_id.CodgenTypeId
	Base() *CodeGenTypeBase
}

type CodeGenTypeBase struct {
	Id          o.Option[string] `json:"id,omitempty" yaml:"id,omitempty"`
	Name        o.Option[string] `json:"name,omitempty" yaml:"name,omitempty"`
	Description o.Option[string] `json:"description,omitempty" yaml:"description,omitempty"`
	Imported    o.Option[bool]   `json:"imported,omitempty" yaml:"imported,omitempty"`
	Public      o.Option[bool]   `json:"public,omitempty" yaml:"public,omitempty"`
	Required    o.Option[bool]   `json:"required,omitempty" yaml:"required,omitempty"`
	Deprecated  o.Option[bool]   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
}

func (t CodeGenTypeBase) HasValidation() bool {
	return false
}

func (t *CodeGenTypeBase) Base() *CodeGenTypeBase {
	return t
}

func (t *CodeGenTypeBase) Merge(o CodeGenTypeBase) bool {
	t.Description = o.Description
	t.Id = o.Id
	t.Imported = o.Imported
	t.Name = o.Name
	t.Required = o.Required
	t.Public = o.Public

	return false
}

func (t CodeGenTypeBase) SchemaId() o.Option[string] {
	return t.Id
}
