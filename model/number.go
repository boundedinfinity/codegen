package model

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// NumberRange Type
//////////////////////////////////////////////////////////////////

type NumberRange[T numberType] struct {
	Min          optioner.Option[T] `json:"min,omitempty"`
	ExclusiveMin optioner.Option[T] `json:"exclusive-min,omitempty"`
	Max          optioner.Option[T] `json:"max,omitempty"`
	ExclusiveMax optioner.Option[T] `json:"exclusive-max,omitempty"`
}

func (t *NumberRange[T]) WithMin(v T) *NumberRange[T] {
	t.Min = optioner.OfZero(v)
	return t
}

func (t *NumberRange[T]) WithMax(v T) *NumberRange[T] {
	t.Max = optioner.OfZero(v)
	return t
}

func (t *NumberRange[T]) WithExclusiveMin(v T) *NumberRange[T] {
	t.ExclusiveMin = optioner.OfZero(v)
	return t
}

func (t *NumberRange[T]) WithExclusiveMax(v T) *NumberRange[T] {
	t.ExclusiveMax = optioner.OfZero(v)
	return t
}

///////////////////////////////////////////////////////////////////
// number Type
//////////////////////////////////////////////////////////////////

type numberType interface {
	int | float64
}

type number[T numberType] struct {
	CodeGenCommon
	MultipleOf optioner.Option[T]                `json:"multiple-of,omitempty"`
	Ranges     optioner.Option[[]NumberRange[T]] `json:"ranges,omitempty"`
}

///////////////////////////////////////////////////////////////////
// number Builders
//////////////////////////////////////////////////////////////////

func (t *number[T]) WithName(v string) *number[T] {
	t.CodeGenCommon.WithName(v)
	return t
}

func (t *number[T]) WithDescription(v string) *number[T] {
	t.CodeGenCommon.WithDescription(v)
	return t
}

func (t *number[T]) WithRequired(v bool) *number[T] {
	t.CodeGenCommon.WithRequired(v)
	return t
}

func (t *number[T]) WithRanges(v ...NumberRange[T]) *number[T] {
	t.Ranges = optioner.OfSlice(v)
	return t
}

///////////////////////////////////////////////////////////////////
// Integer type
//////////////////////////////////////////////////////////////////

type CodeGenInteger number[int]

func (t CodeGenInteger) TypeId() string {
	return "integer"
}

var _ CodeGenType = &CodeGenInteger{}

func (t *CodeGenInteger) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId         string `json:"type-id"`
		CodeGenInteger `json:",inline"`
	}{
		TypeId:         t.TypeId(),
		CodeGenInteger: *t,
	}

	return json.Marshal(dto)
}

func NewInteger() *CodeGenInteger {
	return &CodeGenInteger{}
}

///////////////////////////////////////////////////////////////////
// Float type
//////////////////////////////////////////////////////////////////

type CodeGenFloat number[float64]

func (t CodeGenFloat) TypeId() string {
	return "float"
}

var _ CodeGenType = &CodeGenFloat{}

func (t *CodeGenFloat) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId       string `json:"type-id"`
		CodeGenFloat `json:",inline"`
	}{
		TypeId:       t.TypeId(),
		CodeGenFloat: *t,
	}

	return json.Marshal(dto)
}

func NewFloat() *CodeGenFloat {
	return &CodeGenFloat{}
}
