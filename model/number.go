package model

import (
	"boundedinfinity/codegen/model/type_id"

	"github.com/boundedinfinity/go-commoner/optioner"
)

type number[T int | float64] struct {
	Common
	Min        optioner.Option[T] `json:"min,omitempty"`
	Max        optioner.Option[T] `json:"max,omitempty"`
	MultipleOf optioner.Option[T] `json:"multiple-of,omitempty"`
}

type numberBuilder[T int | float64] struct {
	t *number[T]
}

func buildNumber[T int | float64]() *numberBuilder[T] {
	return &numberBuilder[T]{}
}

func (b *numberBuilder[T]) Done() number[T] {
	return *b.t
}

func (b *numberBuilder[T]) Min(v T) *numberBuilder[T] {
	b.t.Min = optioner.Some(v)
	return b
}

func (b *numberBuilder[T]) Max(v T) *numberBuilder[T] {
	b.t.Max = optioner.Some(v)
	return b
}

func (b *numberBuilder[T]) MultipleOf(v T) *numberBuilder[T] {
	b.t.MultipleOf = optioner.Some(v)
	return b
}

type Int number[int]

func (t Int) TypeId() type_id.TypeId {
	return type_id.Integer
}

var _ Type = &Int{}

func BuildInt() *numberBuilder[int] {
	return buildNumber[int]()
}

type Float number[float64]

func (t Float) TypeId() type_id.TypeId {
	return type_id.Float
}

var _ Type = &Float{}

func BuildFloat() *numberBuilder[float64] {
	return buildNumber[float64]()
}
