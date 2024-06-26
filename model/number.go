package model

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// Integer
///////////////////////////////////////////////////////////////////

func NewInteger() *CodeGenInteger {
	return &CodeGenInteger{}
}

type CodeGenInteger struct {
	number[int]
}

var _ CodeGenType = &CodeGenInteger{}

func (t CodeGenInteger) BaseType() string {
	return "integer"
}

func (t *CodeGenInteger) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId         string `json:"base-type"`
		CodeGenInteger `json:",inline"`
	}{
		TypeId:         t.BaseType(),
		CodeGenInteger: *t,
	}

	return marshalCodeGenType(dto)
}

func (t *CodeGenInteger) WithQName(v string) *CodeGenInteger {
	t.withQName(v)
	return t
}

func (t *CodeGenInteger) WithName(v string) *CodeGenInteger {
	t.withName(v)
	return t
}

func (t *CodeGenInteger) WithPackage(v string) *CodeGenInteger {
	t.withPackage(v)
	return t
}

func (t *CodeGenInteger) WithDescription(v string) *CodeGenInteger {
	t.withDescription(v)
	return t
}

func (t *CodeGenInteger) WithRequired(v bool) *CodeGenInteger {
	t.withRequired(v)
	return t
}

func (t *CodeGenInteger) WithMultipleOf(v int) *CodeGenInteger {
	t.withMultipleOf(v)
	return t
}

func (t *CodeGenInteger) WithRange(v *NumberRange[int]) *CodeGenInteger {
	t.withRange(v)
	return t
}

func (t *CodeGenInteger) WithList(elems ...int) *CodeGenInteger {
	t.withList(elems...)
	return t
}

///////////////////////////////////////////////////////////////////
// Float
///////////////////////////////////////////////////////////////////

func NewFloat() *CodeGenFloat {
	return &CodeGenFloat{}
}

type CodeGenFloat struct {
	number[float64]
}

var _ CodeGenType = &CodeGenFloat{}

func (t CodeGenFloat) BaseType() string {
	return "float"
}

func (t *CodeGenFloat) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId       string `json:"base-type"`
		CodeGenFloat `json:",inline"`
	}{
		TypeId:       t.BaseType(),
		CodeGenFloat: *t,
	}

	return marshalCodeGenType(dto)
}

func (t *CodeGenFloat) WithQName(v string) *CodeGenFloat {
	t.codeGenCommon.withQName(v)
	return t
}

func (t *CodeGenFloat) WithName(v string) *CodeGenFloat {
	t.codeGenCommon.withName(v)
	return t
}

func (t *CodeGenFloat) WithPackage(v string) *CodeGenFloat {
	t.codeGenCommon.withPackage(v)
	return t
}

func (t *CodeGenFloat) WithDescription(v string) *CodeGenFloat {
	t.codeGenCommon.withDescription(v)
	return t
}

func (t *CodeGenFloat) WithRequired(v bool) *CodeGenFloat {
	t.codeGenCommon.withRequired(v)
	return t
}

func (t *CodeGenFloat) WithMultipleOf(v float64) *CodeGenFloat {
	t.MultipleOf = optioner.Some(v)
	return t
}

func (t *CodeGenFloat) WithRange(v NumberRange[float64]) *CodeGenFloat {
	t.Ranges = append(t.Ranges, v)
	return t
}

///////////////////////////////////////////////////////////////////
// number
///////////////////////////////////////////////////////////////////

type number[T int | float64] struct {
	codeGenCommon
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
	if err := t.codeGenCommon.Validate(); err != nil {
		return err
	}

	if t.MultipleOf.Defined() && t.MultipleOf.Get() < 1 {
		return ErrNumberMultipleOfBelow1.WithValue(t.MultipleOf.Get())
	}

	for _, rng := range t.Ranges {
		if err := rng.Validate(); err != nil {
			return err
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

func (t *number[T]) withName(v string) *number[T] {
	t.codeGenCommon.withName(v)
	return t
}

func (t *number[T]) withList(elems ...T) *number[T] {
	for _, elem := range elems {
		t.withRange(NewRange[T]().WithMin(elem).WithMax(elem))
	}
	return t
}

func (t *number[T]) withMultipleOf(v T) *number[T] {
	t.MultipleOf = optioner.Some(v)
	return t
}

func (t *number[T]) withRange(v *NumberRange[T]) *number[T] {
	t.Ranges = append(t.Ranges, *v)
	return t
}

///////////////////////////////////////////////////////////////////
// NumberRange
///////////////////////////////////////////////////////////////////

func NewRange[T int | float64]() *NumberRange[T] {
	return &NumberRange[T]{}
}

type NumberRange[T int | float64] struct {
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
// Marshal
//----------------------------------------------------------------

func (t *NumberRange[T]) MarshalJSON() ([]byte, error) {
	type internal NumberRange[T]
	return marshalCodeGenType(internal(*t))
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

var (
	ErrNumberRangeMinAndExclusiveMinMutuallyExclusive = errorer.New("min and exclusive-min are multually exclusive")
	ErrNumberRangeMinOrExclusiveMinRequired           = errorer.New("min or exclusive-min required")
	ErrNumberRangeMaxAndExclusiveMaxMutuallyExclusive = errorer.New("max and exclusive-max are multually exclusive")
	ErrNumberRangeMaxOrExclusiveMaxRequired           = errorer.New("max or exclusive-max required")
	ErrNumberRangeMaxLessThanMin                      = errorer.New("max less than min")
)

func (t NumberRange[T]) Validate() error {
	switch {
	case t.Min.Defined() && t.ExclusiveMin.Defined():
		return ErrNumberRangeMinAndExclusiveMinMutuallyExclusive
	case t.Min.Empty() && t.ExclusiveMin.Empty():
		return ErrNumberRangeMinOrExclusiveMinRequired
	case t.Max.Defined() && t.ExclusiveMax.Defined():
		return ErrNumberRangeMaxAndExclusiveMaxMutuallyExclusive
	case t.Max.Empty() && t.ExclusiveMax.Empty():
		return ErrNumberRangeMaxOrExclusiveMaxRequired
	}

	switch {
	case t.Max.Defined(), t.Min.Defined():
		if t.Min.Get() > t.Max.Get() {
			return ErrNumberRangeMaxLessThanMin.
				FormatFn("max %d, min %d")(t.Max.Get(), t.Min.Get())
		}
	case t.Max.Defined(), t.ExclusiveMin.Defined():
		if t.Min.Get() >= t.ExclusiveMax.Get() {
			return ErrNumberRangeMaxLessThanMin.
				FormatFn("max %d, min %d")(t.Max.Get(), t.ExclusiveMin.Get())
		}
	case t.ExclusiveMax.Defined(), t.Min.Defined():
		if t.ExclusiveMin.Get() > t.ExclusiveMax.Get() {
			return ErrNumberRangeMaxLessThanMin.
				FormatFn("max %d, min %d")(t.ExclusiveMax.Get(), t.Min.Get())
		}
	case t.ExclusiveMax.Defined(), t.ExclusiveMin.Defined():
		if t.ExclusiveMin.Get() < t.ExclusiveMax.Get() {
			return ErrNumberRangeMaxLessThanMin.
				FormatFn("max %d, min %d")(t.ExclusiveMax.Get(), t.ExclusiveMin.Get())
		}
	}

	return nil
}

func (t NumberRange[T]) HasValidation() bool {
	return t.Max.Defined() || t.ExclusiveMax.Defined() ||
		t.Min.Defined() || t.ExclusiveMin.Defined()
}
