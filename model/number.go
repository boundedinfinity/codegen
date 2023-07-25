package model

import (
	"boundedinfinity/codegen/model/type_id"

	"github.com/boundedinfinity/go-commoner/optioner"
)

// Int

type Int number[int]

func (t Int) TypeId() type_id.TypeId {
	return type_id.Integer
}

var _ Type = &Int{}

func BuildInt() *numberBuilder[int] {
	return buildNumber[int]()
}

// Float

type Float number[float64]

func (t Float) TypeId() type_id.TypeId {
	return type_id.Float
}

var _ Type = &Float{}

func BuildFloat() *numberBuilder[float64] {
	return buildNumber[float64]()
}

// Number

type number[T int | float64] struct {
	Common
	Min        optioner.Option[T] `json:"min,omitempty"`
	Max        optioner.Option[T] `json:"max,omitempty"`
	MultipleOf optioner.Option[T] `json:"multiple-of,omitempty"`
}

// Builder

type numberBuilder[T int | float64] struct {
	t number[T]
}

func buildNumber[T int | float64]() *numberBuilder[T] {
	return &numberBuilder[T]{}
}

func (b *numberBuilder[T]) Done() number[T] {
	return b.t
}

func (b *numberBuilder[T]) Common(v Common) *numberBuilder[T] {
	b.t.Common = v
	return b
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

// Helpers

func validateMinMax[T ~int | ~int64 | ~float64](s string, min, max optioner.Option[T]) error {
	if min.Defined() && max.Defined() {
		if max.Get() < min.Get() {
			return ErrMinMax
		}
	}

	return nil
}
