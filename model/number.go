package model

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Integer
///////////////////////////////////////////////////////////////////

type CodeGenInteger number[int]

func (t CodeGenInteger) BaseType() string {
	return "integer"
}

var _ CodeGenType = &CodeGenInteger{}

func (t *CodeGenInteger) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId         string `json:"base-type"`
		CodeGenInteger `json:",inline"`
	}{
		TypeId:         t.BaseType(),
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

func (t CodeGenFloat) BaseType() string {
	return "float"
}

var _ CodeGenType = &CodeGenFloat{}

func (t *CodeGenFloat) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId       string `json:"base-type"`
		CodeGenFloat `json:",inline"`
	}{
		TypeId:       t.BaseType(),
		CodeGenFloat: *t,
	}

	return json.Marshal(dto)
}

func NewFloat() *CodeGenFloat {
	return &CodeGenFloat{}
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

var (
	ErrNumberMultipleOfBelow1 = errorer.New("multiple of below 1")
	ErrNumberRange            = errorer.New("number range")
)

func (t number[T]) Validate() error {
	if err := t.CodeGenCommon.Validate(); err != nil {
		return err
	}

	if t.MultipleOf.Defined() && t.MultipleOf.Get() < 1 {
		return ErrNumberMultipleOfBelow1.WithValue(t.MultipleOf.Get())
	}

	for i, rng := range t.Ranges {
		if err := rng.Validate(); err != nil {
			return ErrNumberRange.FormatFn("%s [%i]")(err.Error(), i)
		}
	}

	return nil
}

func (t number[T]) HasValidation() bool {
	return t.Common().HasValidation() || t.MultipleOf.Defined() || len(t.Ranges) > 0
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func (t *number[T]) WithQName(v string) *number[T] {
	t.CodeGenCommon.WithQName(v)
	return t
}

func (t *number[T]) WithName(v string) *number[T] {
	t.CodeGenCommon.WithName(v)
	return t
}

func (t *number[T]) WithPackage(v string) *number[T] {
	t.CodeGenCommon.WithPackage(v)
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
// NumberRange
///////////////////////////////////////////////////////////////////

type NumberRange[T numberType] struct {
	Min          optioner.Option[T] `json:"min,omitempty"`
	ExclusiveMin optioner.Option[T] `json:"exclusive-min,omitempty"`
	Max          optioner.Option[T] `json:"max,omitempty"`
	ExclusiveMax optioner.Option[T] `json:"exclusive-max,omitempty"`
}

func (t *NumberRange[T]) WithMin(v T) *NumberRange[T] {
	t.Min = optioner.Some(v)
	return t
}

func (t *NumberRange[T]) WithMax(v T) *NumberRange[T] {
	t.Max = optioner.Some(v)
	return t
}

func (t *NumberRange[T]) WithExclusiveMin(v T) *NumberRange[T] {
	t.ExclusiveMin = optioner.Some(v)
	return t
}

func (t *NumberRange[T]) WithExclusiveMax(v T) *NumberRange[T] {
	t.ExclusiveMax = optioner.Some(v)
	return t
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

var (
	ErrNumberRangeMinAndExclusiveMinMutuallyExclusive = errorer.New("min and exclusive-min are multually exclusive")
	ErrNumberRangeMaxAndExclusiveMaxMutuallyExclusive = errorer.New("max and exclusive-max are multually exclusive")
	ErrNumberRangeMaxLessThanMin                      = errorer.New("max less than min")
)

func (t NumberRange[T]) Validate() error {
	if t.Min.Defined() && t.ExclusiveMin.Defined() {
		return ErrNumberRangeMinAndExclusiveMinMutuallyExclusive
	}

	if t.Max.Defined() && t.ExclusiveMax.Defined() {
		return ErrNumberRangeMaxAndExclusiveMaxMutuallyExclusive
	}

	switch {
	case t.Max.Defined(), t.Min.Defined():
		if t.Min.Get() < t.Max.Get() {
			return ErrNumberRangeMaxLessThanMin.
				FormatFn("max %s, min %s")(t.Max.Get(), t.Min.Get())
		}
	case t.Max.Defined(), t.ExclusiveMin.Defined():
		if t.Min.Get() < t.ExclusiveMax.Get() {
			return ErrNumberRangeMaxLessThanMin.
				FormatFn("max %s, min %s")(t.Max.Get(), t.ExclusiveMin.Get())
		}
	case t.ExclusiveMax.Defined(), t.Min.Defined():
		if t.Min.Get() < t.ExclusiveMax.Get() {
			return ErrNumberRangeMaxLessThanMin.
				FormatFn("max %s, min %s")(t.ExclusiveMax.Get(), t.Min.Get())
		}
	case t.ExclusiveMax.Defined(), t.ExclusiveMin.Defined():
		if t.Min.Get() < t.ExclusiveMax.Get() {
			return ErrNumberRangeMaxLessThanMin.
				FormatFn("max %s, min %s")(t.ExclusiveMax.Get(), t.ExclusiveMin.Get())
		}
	}

	return nil
}

func (t NumberRange[T]) HasValidation() bool {
	return t.Max.Defined() || t.ExclusiveMax.Defined() ||
		t.Min.Defined() || t.ExclusiveMin.Defined()
}
