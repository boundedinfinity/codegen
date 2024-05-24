package model

import "github.com/boundedinfinity/go-commoner/functional/optioner"

type number[T int | float64] struct {
	CodeGenCommon
	Min        optioner.Option[T] `json:"min,omitempty"`
	Max        optioner.Option[T] `json:"max,omitempty"`
	MultipleOf optioner.Option[T] `json:"multiple-of,omitempty"`
}

type CodeGenInteger number[int]

func (t CodeGenInteger) TypeId() string {
	return "integer"
}

var _ CodeGenType = &CodeGenInteger{}

type CodeGenFloat number[float64]

func (t CodeGenFloat) TypeId() string {
	return "float"
}

var _ CodeGenType = &CodeGenFloat{}
