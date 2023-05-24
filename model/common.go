package model

import (
	"boundedinfinity/codegen/model/type_id"

	"github.com/boundedinfinity/go-commoner/optioner"
)

type Type interface {
	TypeId() type_id.TypeId
}

type common struct {
	Name        optioner.Option[string] `json:"name,omitempty"`
	Description optioner.Option[string] `json:"description,omitempty"`
	Required    optioner.Option[bool]   `json:"required,omitempty"`
	// Default     optioner.Option[T]      `json:"default,omitempty"`
}
