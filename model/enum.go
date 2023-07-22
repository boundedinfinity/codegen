package model

import (
	"boundedinfinity/codegen/model/type_id"
)

type Enum struct {
	Common
	Values []string
}

func (t Enum) TypeId() type_id.TypeId {
	return type_id.Enum
}

var _ Type = &Enum{}

type enumBuilder struct {
	t *Enum
}

func BuildEnum() *enumBuilder {
	return &enumBuilder{}
}

func (b *enumBuilder) Done() Enum {
	return *b.t
}

func (b *enumBuilder) Values(v ...string) *enumBuilder {
	b.t.Values = v
	return b
}
