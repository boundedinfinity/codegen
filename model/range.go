package model

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

///////////////////////////////////////////////////////////////////
// NumberRange
///////////////////////////////////////////////////////////////////

func NewRange[T NumberType]() *Range[T] {
	return &Range[T]{}
}

type Range[T NumberType] struct {
	Min          optioner.Option[T] `json:"min,omitempty"`
	ExclusiveMin optioner.Option[T] `json:"exclusive-min,omitempty"`
	Max          optioner.Option[T] `json:"max,omitempty"`
	ExclusiveMax optioner.Option[T] `json:"exclusive-max,omitempty"`
}

type RangeBuilder[T NumberType] struct {
	obj Range[T]
}

func BuildRange[T ~int | ~float64]() *RangeBuilder[T] {
	return &RangeBuilder[T]{}
}

func (this *RangeBuilder[T]) Build() Range[T] {
	return this.obj
}

func (this *RangeBuilder[T]) Min(v T) *RangeBuilder[T] {
	return SetO(this, &this.obj.Min, v)
}

func (this *RangeBuilder[T]) Max(v T) *RangeBuilder[T] {
	return SetO(this, &this.obj.Max, v)
}

func (this *RangeBuilder[T]) ExclusiveMin(v T) *RangeBuilder[T] {
	return SetO(this, &this.obj.ExclusiveMin, v)
}

func (this *RangeBuilder[T]) ExclusiveMax(v T) *RangeBuilder[T] {
	return SetO(this, &this.obj.ExclusiveMax, v)
}

//----------------------------------------------------------------
// Marshal
//----------------------------------------------------------------

func (this *Range[T]) MarshalJSON() ([]byte, error) {
	type internal Range[T]
	return marshalCodeGenType(internal(*this))
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

type errErrRangeMinAndExclusiveMinMutuallyExclusive[T NumberType] struct {
	Min          T
	ExclusiveMin T
}

func (e errErrRangeMinAndExclusiveMinMutuallyExclusive[T]) Error() string {
	return fmt.Sprintf("min: %v, exclusive-min: %v : %s", e.Min, e.ExclusiveMin,
		ErrNumberRangeMinAndExclusiveMinMutuallyExclusive.Error())
}

func (e errErrRangeMinAndExclusiveMinMutuallyExclusive[T]) Unwrap() error {
	return ErrNumberRangeMinAndExclusiveMinMutuallyExclusive
}

type errErrRangeMaxAndExclusiveMaxMutuallyExclusive[T NumberType] struct {
	Max          T
	ExclusiveMax T
}

func (e errErrRangeMaxAndExclusiveMaxMutuallyExclusive[T]) Error() string {
	return fmt.Sprintf("max: %v, exclusive-max: %v : %s", e.Max, e.ExclusiveMax,
		ErrNumberRangeMaxAndExclusiveMaxMutuallyExclusive.Error())
}

func (e errErrRangeMaxAndExclusiveMaxMutuallyExclusive[T]) Unwrap() error {
	return ErrNumberRangeMinAndExclusiveMinMutuallyExclusive
}

type errRangeMinMax[T NumberType] struct {
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

func (this Range[T]) Validate() error {
	switch {
	case this.Min.Defined() && this.ExclusiveMin.Defined():
		return &errErrRangeMinAndExclusiveMinMutuallyExclusive[T]{this.Min.Get(), this.ExclusiveMin.Get()}
	case this.Min.Empty() && this.ExclusiveMin.Empty():
		return ErrNumberRangeMinOrExclusiveMinRequired
	case this.Max.Defined() && this.ExclusiveMax.Defined():
		return &errErrRangeMaxAndExclusiveMaxMutuallyExclusive[T]{this.Max.Get(), this.ExclusiveMax.Get()}
	case this.Max.Empty() && this.ExclusiveMax.Empty():
		return ErrNumberRangeMaxOrExclusiveMaxRequired
	}

	switch {
	case this.Max.Defined(), this.Min.Defined():
		if this.Min.Get() > this.Max.Get() {
			return &errRangeMinMax[T]{this.Min.Get(), this.Max.Get(), ErrRangeMaxLessThanMin}
		}
	case this.Max.Defined(), this.ExclusiveMin.Defined():
		if this.Min.Get() >= this.ExclusiveMax.Get() {
			return &errRangeMinMax[T]{this.ExclusiveMin.Get(), this.Max.Get(), ErrRangeMaxLessThanMin}
		}
	case this.ExclusiveMax.Defined(), this.Min.Defined():
		if this.ExclusiveMin.Get() > this.ExclusiveMax.Get() {
			return &errRangeMinMax[T]{this.Min.Get(), this.ExclusiveMax.Get(), ErrRangeMaxLessThanMin}
		}
	case this.ExclusiveMax.Defined(), this.ExclusiveMin.Defined():
		if this.ExclusiveMin.Get() < this.ExclusiveMax.Get() {
			return &errRangeMinMax[T]{this.ExclusiveMin.Get(), this.ExclusiveMax.Get(), ErrRangeMaxLessThanMin}
		}
	}

	return nil
}

func (this Range[T]) HasValidation() bool {
	return this.Max.Defined() || this.ExclusiveMax.Defined() ||
		this.Min.Defined() || this.ExclusiveMin.Defined()
}
