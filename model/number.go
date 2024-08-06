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
	Number[int]
}

var _ CodeGenSchema = &CodeGenInteger{}

func (t CodeGenInteger) Schema() string {
	return "integer"
}

func (t *CodeGenInteger) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId         string `json:"type"`
		CodeGenInteger `json:",inline"`
	}{
		TypeId:         t.Schema(),
		CodeGenInteger: *t,
	}

	return marshalCodeGenType(dto)
}

///////////////////////////////////////////////////////////////////
// Float
///////////////////////////////////////////////////////////////////

func NewFloat() *CodeGenFloat {
	return &CodeGenFloat{}
}

type CodeGenFloat struct {
	Number[float64]
	Precision optioner.Option[int]     `json:"precision,omitempty"`
	Tolerance optioner.Option[float64] `json:"tolerance,omitempty"`
}

var _ CodeGenSchema = &CodeGenFloat{}

func (t CodeGenFloat) Schema() string {
	return "float"
}

func (t *CodeGenFloat) MarshalJSON() ([]byte, error) {
	dto := struct {
		TypeId       string `json:"type"`
		CodeGenFloat `json:",inline"`
	}{
		TypeId:       t.Schema(),
		CodeGenFloat: *t,
	}

	return marshalCodeGenType(dto)
}

///////////////////////////////////////////////////////////////////
// number
///////////////////////////////////////////////////////////////////

type Number[T ~int | ~float64] struct {
	CodeGenCommon
	Min        optioner.Option[T]                `json:"min,omitempty"`
	Max        optioner.Option[T]                `json:"max,omitempty"`
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

func (t Number[T]) Validate() error {
	if err := t.CodeGenCommon.Validate(); err != nil {
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

func (t Number[T]) HasValidation() bool {
	return t.CodeGenCommon.HasValidation() || t.MultipleOf.Defined() || t.Ranges.Defined()
}

///////////////////////////////////////////////////////////////////
// NumberRange
///////////////////////////////////////////////////////////////////

func NewRange[T ~int | ~float64]() *NumberRange[T] {
	return &NumberRange[T]{}
}

type NumberRange[T ~int | ~float64] struct {
	Min          optioner.Option[T] `json:"min,omitempty"`
	ExclusiveMin optioner.Option[T] `json:"exclusive-min,omitempty"`
	Max          optioner.Option[T] `json:"max,omitempty"`
	ExclusiveMax optioner.Option[T] `json:"exclusive-max,omitempty"`
}

type NumberRangeBuilder[T ~int | ~float64] struct {
	obj NumberRange[T]
}

func BuildRange[T ~int | ~float64]() *NumberRangeBuilder[T] {
	return &NumberRangeBuilder[T]{}
}

func (t *NumberRangeBuilder[T]) Build() NumberRange[T] {
	return t.obj
}

func (t *NumberRangeBuilder[T]) Min(v T) *NumberRangeBuilder[T] {
	return SetO(t, &t.obj.Min, v)
}

func (t *NumberRangeBuilder[T]) Max(v T) *NumberRangeBuilder[T] {
	return SetO(t, &t.obj.Max, v)
}

func (t *NumberRangeBuilder[T]) ExclusiveMin(v T) *NumberRangeBuilder[T] {
	return SetO(t, &t.obj.ExclusiveMin, v)
}

func (t *NumberRangeBuilder[T]) ExclusiveMax(v T) *NumberRangeBuilder[T] {
	return SetO(t, &t.obj.ExclusiveMax, v)
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
