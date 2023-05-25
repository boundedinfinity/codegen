package model

import (
	"boundedinfinity/codegen/model/type_id"

	"github.com/boundedinfinity/go-commoner/optioner"
)

type Type interface {
	TypeId() type_id.TypeId
}

type Common struct {
	Name        optioner.Option[string] `json:"name,omitempty"`
	Description optioner.Option[string] `json:"description,omitempty"`
	Required    optioner.Option[bool]   `json:"required,omitempty"`
	// Default     optioner.Option[T]      `json:"default,omitempty"`
}

type ParamT interface {
	Common | String
}

type ParamFunc[T ParamT] func(t *T)

func handleParams[T ParamT](t *T, params ...ParamFunc[T]) {
	for _, param := range params {
		param(t)
	}
}

func Name(v string) func(t *Common) {
	return NameOf(optioner.Some(v))
}

func NameOf(v optioner.Option[string]) func(t *Common) {
	return func(t *Common) {
		t.Name = v
	}
}

func Description(v string) func(t *Common) {
	return func(t *Common) {
		t.Description = optioner.Some(v)
	}
}

func Required(v bool) func(t *Common) {
	return func(t *Common) {
		t.Required = optioner.Some(v)
	}
}
