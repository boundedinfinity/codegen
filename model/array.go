package model

import (
	"boundedinfinity/codegen/model/type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type Array struct {
	Common
	Items Type
	Min   o.Option[int] `json:"min,omitempty"`
	Max   o.Option[int] `json:"max,omitempty"`
}

func (t Array) TypeId() type_id.TypeId {
	return type_id.Array
}

var _ Type = &Array{}

type arrayBuilder struct {
	t Array
}

func BuildArray() *arrayBuilder {
	return &arrayBuilder{}
}

func (b *arrayBuilder) Done() Array {
	return b.t
}

func (b *arrayBuilder) Min(v int) *arrayBuilder {
	b.t.Min = o.Some(v)
	return b
}

func (b *arrayBuilder) Max(v int) *arrayBuilder {
	b.t.Max = o.Some(v)
	return b
}
