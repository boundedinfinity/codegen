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

type Int number[int]

func (t Int) TypeId() type_id.TypeId {
	return type_id.Integer
}

var _ Type = &Int{}

type Float number[float64]

func (t Float) TypeId() type_id.TypeId {
	return type_id.Float
}

var _ Type = &Float{}
