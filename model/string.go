package model

import (
	"boundedinfinity/codegen/model/type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type String struct {
	Common
	Min   o.Option[int]    `json:"min,omitempty"`
	Max   o.Option[int]    `json:"max,omitempty"`
	Regex o.Option[string] `json:"regex,omitempty"`
}

func (t String) TypeId() type_id.TypeId {
	return type_id.String
}

var _ Type = &String{}

type stringBuilder struct {
	t *String
}

func BuildString() *stringBuilder {
	return &stringBuilder{}
}

func (b *stringBuilder) Done() String {
	return *b.t
}

func (b *stringBuilder) Min(v int) *stringBuilder {
	b.t.Min = o.Some(v)
	return b
}

func (b *stringBuilder) Max(v int) *stringBuilder {
	b.t.Max = o.Some(v)
	return b
}

func (b *stringBuilder) Regex(v string) *stringBuilder {
	b.t.Regex = o.Some(v)
	return b
}
