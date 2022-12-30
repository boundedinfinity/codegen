package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenType interface {
	SchemaType() codegen_type_id.CodgenTypeId
	Base() *CodeGenTypeBase
	HasValidation() bool
	ValidateSchema() error
}

type CodeGenTypeBase struct {
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
