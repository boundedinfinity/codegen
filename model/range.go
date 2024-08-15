package model

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// NumberRange
///////////////////////////////////////////////////////////////////

func NewRange[T ~int | ~float64]() *Range[T] {
	return &Range[T]{}
}

type Range[T ~int | ~float64] struct {
	Min          optioner.Option[T] `json:"min,omitempty"`
	ExclusiveMin optioner.Option[T] `json:"exclusive-min,omitempty"`
	Max          optioner.Option[T] `json:"max,omitempty"`
	ExclusiveMax optioner.Option[T] `json:"exclusive-max,omitempty"`
}

type RangeBuilder[T ~int | ~float64] struct {
	obj Range[T]
}

func BuildRange[T ~int | ~float64]() *RangeBuilder[T] {
	return &RangeBuilder[T]{}
}

func (t *RangeBuilder[T]) Build() Range[T] {
	return t.obj
}

func (t *RangeBuilder[T]) Min(v T) *RangeBuilder[T] {
	return SetO(t, &t.obj.Min, v)
}

func (t *RangeBuilder[T]) Max(v T) *RangeBuilder[T] {
	return SetO(t, &t.obj.Max, v)
}

func (t *RangeBuilder[T]) ExclusiveMin(v T) *RangeBuilder[T] {
	return SetO(t, &t.obj.ExclusiveMin, v)
}

func (t *RangeBuilder[T]) ExclusiveMax(v T) *RangeBuilder[T] {
	return SetO(t, &t.obj.ExclusiveMax, v)
}

//----------------------------------------------------------------
// Marshal
//----------------------------------------------------------------

func (t *Range[T]) MarshalJSON() ([]byte, error) {
	type internal Range[T]
	return marshalCodeGenType(internal(*t))
}

//----------------------------------------------------------------
// Validation
//----------------------------------------------------------------

var (
	ErrNumberRangeMinAndExclusiveMinMutuallyExclusive = errors.New("min and exclusive-min are multually exclusive")
	ErrNumberRangeMinOrExclusiveMinRequired           = errors.New("min or exclusive-min required")
	ErrNumberRangeMaxAndExclusiveMaxMutuallyExclusive = errors.New("max and exclusive-max are multually exclusive")
	ErrNumberRangeMaxOrExclusiveMaxRequired           = errors.New("max or exclusive-max required")
	ErrRangeMaxLessThanMin                            = errors.New("max less than min")
)

type errErrRangeMinAndExclusiveMinMutuallyExclusive[T ~int | ~float64] struct {
	Min          T
	ExclusiveMin T
}

func (e errErrRangeMinAndExclusiveMinMutuallyExclusive[T]) Error() string {
	return fmt.Sprintf("min: %v, exclusive-min: %v : %s", e.Min, e.ExclusiveMin, ErrNumberRangeMinAndExclusiveMinMutuallyExclusive.Error())
}

func (e errErrRangeMinAndExclusiveMinMutuallyExclusive[T]) Unwrap() error {
	return ErrNumberRangeMinAndExclusiveMinMutuallyExclusive
}

type errErrRangeMaxAndExclusiveMaxMutuallyExclusive[T ~int | ~float64] struct {
	Max          T
	ExclusiveMax T
}

func (e errErrRangeMaxAndExclusiveMaxMutuallyExclusive[T]) Error() string {
	return fmt.Sprintf("max: %v, exclusive-max: %v : %s", e.Max, e.ExclusiveMax, ErrNumberRangeMaxAndExclusiveMaxMutuallyExclusive.Error())
}

func (e errErrRangeMaxAndExclusiveMaxMutuallyExclusive[T]) Unwrap() error {
	return ErrNumberRangeMinAndExclusiveMinMutuallyExclusive
}

type errRangeMinMax[T ~int | ~float64] struct {
	Min    T
	Max    T
	parent error
}

func (e errRangeMinMax[T]) Error() string {
	return fmt.Sprintf("min: %v, max: %v : %s", e.Min, e.Max, e.parent.Error())
}

func (e errRangeMinMax[T]) Unwrap() error {
	return e.parent
}

func (t Range[T]) Validate() error {
	switch {
	case t.Min.Defined() && t.ExclusiveMin.Defined():
		return &errErrRangeMinAndExclusiveMinMutuallyExclusive[T]{t.Min.Get(), t.ExclusiveMin.Get()}
	case t.Min.Empty() && t.ExclusiveMin.Empty():
		return ErrNumberRangeMinOrExclusiveMinRequired
	case t.Max.Defined() && t.ExclusiveMax.Defined():
		return &errErrRangeMaxAndExclusiveMaxMutuallyExclusive[T]{t.Max.Get(), t.ExclusiveMax.Get()}
	case t.Max.Empty() && t.ExclusiveMax.Empty():
		return ErrNumberRangeMaxOrExclusiveMaxRequired
	}

	switch {
	case t.Max.Defined(), t.Min.Defined():
		if t.Min.Get() > t.Max.Get() {
			return &errRangeMinMax[T]{t.Min.Get(), t.Max.Get(), ErrRangeMaxLessThanMin}
		}
	case t.Max.Defined(), t.ExclusiveMin.Defined():
		if t.Min.Get() >= t.ExclusiveMax.Get() {
			return &errRangeMinMax[T]{t.ExclusiveMin.Get(), t.Max.Get(), ErrRangeMaxLessThanMin}
		}
	case t.ExclusiveMax.Defined(), t.Min.Defined():
		if t.ExclusiveMin.Get() > t.ExclusiveMax.Get() {
			return &errRangeMinMax[T]{t.Min.Get(), t.ExclusiveMax.Get(), ErrRangeMaxLessThanMin}
		}
	case t.ExclusiveMax.Defined(), t.ExclusiveMin.Defined():
		if t.ExclusiveMin.Get() < t.ExclusiveMax.Get() {
			return &errRangeMinMax[T]{t.ExclusiveMin.Get(), t.ExclusiveMax.Get(), ErrRangeMaxLessThanMin}
		}
	}

	return nil
}

func (t Range[T]) HasValidation() bool {
	return t.Max.Defined() || t.ExclusiveMax.Defined() ||
		t.Min.Defined() || t.ExclusiveMin.Defined()
}
