package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// NumberRange
///////////////////////////////////////////////////////////////////

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

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t *NumberRange[T]) Validate() error {
	if t.Min.Defined() && t.ExclusiveMin.Defined() {
		return errors.New("min and exclusive-min are multually exclusive")
	}

	if t.Max.Defined() && t.ExclusiveMax.Defined() {
		return errors.New("max and exclusive-max are multually exclusive")
	}

	return nil
}

///////////////////////////////////////////////////////////////////
// numberType
///////////////////////////////////////////////////////////////////

type numberType interface {
	int | float64
}

type number[T numberType] struct {
	CodeGenCommon
	MultipleOf optioner.Option[T] `json:"multiple-of,omitempty"`
	Ranges     []NumberRange[T]   `json:"ranges,omitempty"`
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

func (t *number[T]) Validate() error {
	if err := t.CodeGenCommon.Validate(); err != nil {
		return err
	}

	for i, rng := range t.Ranges {
		if err := rng.Validate(); err != nil {
			return errors.Join(fmt.Errorf("range[%v]", i))
		}
	}

	return nil
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func (t *number[T]) WithSchemaId(v string) *number[T] {
	t.CodeGenCommon.WithTypeId(v)
	return t
}

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
	t.Ranges = append(t.Ranges, v...)
	return t
}

///////////////////////////////////////////////////////////////////
// Integer
///////////////////////////////////////////////////////////////////

type CodeGenInteger number[int]

func (t CodeGenInteger) CodeGenId() string {
	return "integer"
}

var _ CodeGenType = &CodeGenInteger{}

func (t *CodeGenInteger) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId         string `json:"codegen-id"`
		CodeGenInteger `json:",inline"`
	}{
		TypeId:         t.CodeGenId(),
		CodeGenInteger: *t,
	}

	return json.Marshal(dto)
}

func NewInteger() *CodeGenInteger {
	return &CodeGenInteger{}
}

///////////////////////////////////////////////////////////////////
// Float
///////////////////////////////////////////////////////////////////

type CodeGenFloat number[float64]

func (t CodeGenFloat) CodeGenId() string {
	return "float"
}

var _ CodeGenType = &CodeGenFloat{}

func (t *CodeGenFloat) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId       string `json:"codegen-id"`
		CodeGenFloat `json:",inline"`
	}{
		TypeId:       t.CodeGenId(),
		CodeGenFloat: *t,
	}

	return json.Marshal(dto)
}

func NewFloat() *CodeGenFloat {
	return &CodeGenFloat{}
}
