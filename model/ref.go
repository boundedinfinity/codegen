package model

import (
	"boundedinfinity/codegen/model/type_id"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type Ref struct {
	Common
	Ref o.Option[string]
}

func (t Ref) TypeId() type_id.TypeId {
	return type_id.Ref
}

var _ Type = &Ref{}

type refBuilder struct {
	t *Ref
}

func BuildRef() *refBuilder {
	return &refBuilder{}
}

func (b *refBuilder) Done() Ref {
	return *b.t
}

func (b *refBuilder) Ref(v string) *refBuilder {
	b.t.Ref = o.Some(v)
	return b
}
