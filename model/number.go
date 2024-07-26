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

func (t CodeGenInteger) GetType() string {
	return "integer"
}

func (t *CodeGenInteger) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId         string `json:"base-type"`
		CodeGenInteger `json:",inline"`
	}{
		TypeId:         t.GetType(),
		CodeGenInteger: *t,
	}

	return marshalCodeGenType(dto)
}

//----------------------------------------------------------------
// Builder
//----------------------------------------------------------------

func BuildInteger() IntBuilder {
	return &codeGenIntBuilder{}
}

type codeGenIntBuilder struct {
	obj CodeGenInteger
}

var _ IntBuilder = &codeGenIntBuilder{}

// Build implements IntBuilder.
func (t *codeGenIntBuilder) Build() *CodeGenInteger {
	return &t.obj
}

// Description implements IntBuilder.
func (t *codeGenIntBuilder) Description(v string) IntBuilder {
	return setO(t, &t.obj.Description, v)
}

// MultipleOf implements IntBuilder.
func (t *codeGenIntBuilder) MultipleOf(v int) IntBuilder {
	return setO(t, &t.obj.MultipleOf, v)
}

// Name implements IntBuilder.
func (t *codeGenIntBuilder) Name(v string) IntBuilder {
	return setO(t, &t.obj.Name, v)
}

// Negative implements IntBuilder.
func (t *codeGenIntBuilder) Negative() IntBuilder {
	return setO(t, &t.obj.Negative, true)
}

// NoneOf implements IntBuilder.
func (t *codeGenIntBuilder) NoneOf(v ...int) IntBuilder {
	return setO(t, &t.obj.NoneOf, v)
}

// OneOf implements IntBuilder.
func (t *codeGenIntBuilder) OneOf(v ...int) IntBuilder {
	return setO(t, &t.obj.OneOf, v)
}

// Package implements IntBuilder.
func (t *codeGenIntBuilder) Package(v string) IntBuilder {
	return setO(t, &t.obj.Package, v)
}

// Positive implements IntBuilder.
func (t *codeGenIntBuilder) Positive() IntBuilder {
	return setO(t, &t.obj.Positive, true)
}

// QName implements IntBuilder.
func (t *codeGenIntBuilder) QName(string) IntBuilder {
	panic("unimplemented")
}

// Ranges implements IntBuilder.
func (t *codeGenIntBuilder) Ranges(v ...NumberRange[int]) IntBuilder {
	return setO(t, &t.obj.Ranges, v)
}

// Ref implements IntBuilder.
func (t *codeGenIntBuilder) Ref() RefBuilder {
	panic("unimplemented")
}

// Required implements IntBuilder.
func (t *codeGenIntBuilder) Required(v bool) IntBuilder {
	return setO(t, &t.obj.Required, v)
}

///////////////////////////////////////////////////////////////////
// Float
///////////////////////////////////////////////////////////////////

func NewFloat() *CodeGenFloat {
	return &CodeGenFloat{}
}

type CodeGenFloat struct {
	number[float64]
	Precision optioner.Option[int]     `json:"precision,omitempty"`
	Tolerance optioner.Option[float64] `json:"tolerance,omitempty"`
}

var _ CodeGenType = &CodeGenFloat{}

func (t CodeGenFloat) GetType() string {
	return "float"
}

func (t *CodeGenFloat) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId       string `json:"base-type"`
		CodeGenFloat `json:",inline"`
	}{
		TypeId:       t.GetType(),
		CodeGenFloat: *t,
	}

	return marshalCodeGenType(dto)
}

//----------------------------------------------------------------
// Builder
//----------------------------------------------------------------

func BuildFloat() FloatBuilder {
	return &codeGenFloatBuilder{}
}

type codeGenFloatBuilder struct {
	obj CodeGenFloat
}

var _ FloatBuilder = &codeGenFloatBuilder{}

// Build implements FloatBuilder.
func (t *codeGenFloatBuilder) Build() *CodeGenFloat {
	return &t.obj
}

// Description implements FloatBuilder.
func (t *codeGenFloatBuilder) Description(v string) FloatBuilder {
	return setO(t, &t.obj.Description, v)
}

// MultipleOf implements FloatBuilder.
func (t *codeGenFloatBuilder) MultipleOf(v float64) FloatBuilder {
	return setO(t, &t.obj.MultipleOf, v)
}

// Name implements FloatBuilder.
func (t *codeGenFloatBuilder) Name(v string) FloatBuilder {
	return setO(t, &t.obj.Name, v)
}

// Negative implements FloatBuilder.
func (t *codeGenFloatBuilder) Negative() FloatBuilder {
	return setO(t, &t.obj.Negative, true)
}

// NoneOf implements FloatBuilder.
func (t *codeGenFloatBuilder) NoneOf(v ...float64) FloatBuilder {
	return setO(t, &t.obj.NoneOf, v)
}

// OneOf implements FloatBuilder.
func (t *codeGenFloatBuilder) OneOf(v ...float64) FloatBuilder {
	return setO(t, &t.obj.OneOf, v)
}

// Package implements FloatBuilder.
func (t *codeGenFloatBuilder) Package(v string) FloatBuilder {
	return setO(t, &t.obj.Package, v)
}

// Positive implements FloatBuilder.
func (t *codeGenFloatBuilder) Positive() FloatBuilder {
	return setO(t, &t.obj.Positive, true)
}

// Precision implements FloatBuilder.
func (t *codeGenFloatBuilder) Precision(v int) FloatBuilder {
	return setO(t, &t.obj.Precision, v)
}

// QName implements FloatBuilder.
func (t *codeGenFloatBuilder) QName(v string) FloatBuilder {
	panic("unimplemented")
}

// Ranges implements FloatBuilder.
func (t *codeGenFloatBuilder) Ranges(v ...NumberRange[float64]) FloatBuilder {
	return setO(t, &t.obj.Ranges, v)
}

// Ref implements FloatBuilder.
func (t *codeGenFloatBuilder) Ref() RefBuilder {
	panic("unimplemented")
}

// Required implements FloatBuilder.
func (t *codeGenFloatBuilder) Required(v bool) FloatBuilder {
	return setO(t, &t.obj.Required, v)
}

// Tolerance implements FloatBuilder.
func (t *codeGenFloatBuilder) Tolerance(v float64) FloatBuilder {
	return setO(t, &t.obj.Tolerance, v)
}

///////////////////////////////////////////////////////////////////
// number
///////////////////////////////////////////////////////////////////

type number[T int | float64] struct {
	codeGenCommon
	MultipleOf optioner.Option[T]                `json:"multiple-of,omitempty"`
	Ranges     optioner.Option[[]NumberRange[T]] `json:"ranges,omitempty"`
	NoneOf     optioner.Option[[]T]              `json:"one-of,omitempty"`
	OneOf      optioner.Option[[]T]              `json:"some-of,omitempty"`
	Positive   optioner.Option[bool]             `json:"positive,omitempty"`
	Negative   optioner.Option[bool]             `json:"negative,omitempty"`
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

var (
	ErrNumberMultipleOfBelow1                          = errorer.New("multiple of below 1")
	ErrNumberRange                                     = errorer.New("number range")
	ErrNumberRangePositiveAndNegativeMutuallyExclusive = errorer.New("positive and negative are multually exclusive")
)

func (t number[T]) Validate() error {
	if err := t.codeGenCommon.Validate(); err != nil {
		return err
	}

	if t.MultipleOf.Defined() && t.MultipleOf.Get() < 1 {
		return ErrNumberMultipleOfBelow1.WithValue(t.MultipleOf.Get())
	}

	if t.Positive.Defined() && t.Negative.Defined() {
		return ErrNumberRangePositiveAndNegativeMutuallyExclusive
	}

	for _, rng := range t.Ranges.Get() {
		if err := rng.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (t number[T]) HasValidation() bool {
	return t.codeGenCommon.HasValidation() || t.MultipleOf.Defined() || t.Ranges.Defined()
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

type NumberRangeBuilder[T int | float64] struct {
	obj NumberRange[T]
}

func BuildRange[T int | float64]() *NumberRangeBuilder[T] {
	return &NumberRangeBuilder[T]{}
}

func (t *NumberRangeBuilder[T]) Build() NumberRange[T] {
	return t.obj
}

func (t *NumberRangeBuilder[T]) Min(v T) *NumberRangeBuilder[T] {
	return setO(t, &t.obj.Min, v)
}

func (t *NumberRangeBuilder[T]) Max(v T) *NumberRangeBuilder[T] {
	return setO(t, &t.obj.Max, v)
}

func (t *NumberRangeBuilder[T]) ExclusiveMin(v T) *NumberRangeBuilder[T] {
	return setO(t, &t.obj.ExclusiveMin, v)
}

func (t *NumberRangeBuilder[T]) ExclusiveMax(v T) *NumberRangeBuilder[T] {
	return setO(t, &t.obj.ExclusiveMax, v)
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
