package validation

import (
	"boundedinfinity/codegen/model"
	"errors"
	"fmt"
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"golang.org/x/exp/constraints"
)

//////////////////////////////////////////////////////////////
// Less than
//////////////////////////////////////////////////////////////

var ErrIntegerMin = errors.New("less than minimum")

type ErrIntegerMinDetails[T constraints.Integer] struct {
	Name string
	Min  T
}

func (e *ErrIntegerMinDetails[T]) Error() string {
	return fmt.Sprintf("%s value of %d is %s", e.Name, e.Min, ErrIntegerMin.Error())
}

func (e *ErrIntegerMinDetails[T]) Unwrap() error {
	return ErrIntegerMin
}

func IntegerMin[T constraints.Integer](name string, min T, value T) error {
	if value < min {
		return &ErrIntegerMinDetails[T]{name, min}
	}

	return nil
}

func IntegerMinFn[T constraints.Integer](name string, min T) func(T) error {
	return func(value T) error { return IntegerMin(name, min, value) }
}

//////////////////////////////////////////////////////////////
// Greater than
//////////////////////////////////////////////////////////////

var ErrIntegerMax = errors.New("is greater than maxium")

type ErrIntegerMaxDetails[T constraints.Integer] struct {
	Name string
	Max  T
}

func (e *ErrIntegerMaxDetails[T]) Error() string {
	return fmt.Sprintf("%s value of %d is %s", e.Name, e.Max, ErrIntegerMax.Error())
}

func (e *ErrIntegerMaxDetails[T]) Unwrap() error {
	return ErrIntegerMax
}

func IntegerMax[T constraints.Integer](name string, max T, value T) error {
	if value > max {
		return &ErrIntegerMaxDetails[T]{Name: name, Max: max}
	}

	return nil
}

func IntegerMaxFn[T constraints.Integer](name string, max T) func(T) error {
	return func(value T) error { return IntegerMax(name, max, value) }
}

//////////////////////////////////////////////////////////////
// Greater than
//////////////////////////////////////////////////////////////

var ErrIntegerRange = errors.New("out of range")

type ErrIntegerRangeDetails[T constraints.Integer] struct {
	Name  string
	Range model.Range[T]
	Value T
}

func (e ErrIntegerRangeDetails[T]) Error() string {
	var sb strings.Builder

	sb.WriteString(e.Name + " is out of range of ")

	switch {
	case e.Range.Min.Defined():
		sb.WriteString(fmt.Sprintf("%d < ", e.Range.Min.Get()))
	case e.Range.ExclusiveMin.Defined():
		sb.WriteString(fmt.Sprintf("%d <= ", e.Range.ExclusiveMin.Get()))
	default:
		panic("min or exclusive min not defined")
	}

	sb.WriteString(fmt.Sprintf("%d", e.Value))

	switch {
	case e.Range.Max.Defined():
		sb.WriteString(fmt.Sprintf(" > %d", e.Range.Min.Get()))
	case e.Range.ExclusiveMax.Defined():
		sb.WriteString(fmt.Sprintf(" >= %d", e.Range.ExclusiveMax.Get()))
	default:
		panic("max or exclusive max not defined")
	}

	return sb.String()
}

func (e *ErrIntegerRangeDetails[T]) Unwrap() error {
	return ErrIntegerRange
}

func IntegerRange[T constraints.Integer](name string, rng model.Range[T]) func(v T) error {
	return func(v T) error {
		if rng.Min.Defined() && v < rng.Min.Get() {
			return &ErrIntegerRangeDetails[T]{Name: name, Range: rng, Value: v}
		}

		if rng.ExclusiveMin.Defined() && v <= rng.ExclusiveMin.Get() {
			return &ErrIntegerRangeDetails[T]{Name: name, Range: rng, Value: v}
		}

		if rng.Max.Defined() && v < rng.Max.Get() {
			return &ErrIntegerRangeDetails[T]{Name: name, Range: rng, Value: v}
		}

		if rng.ExclusiveMax.Defined() && v <= rng.ExclusiveMax.Get() {
			return &ErrIntegerRangeDetails[T]{Name: name, Range: rng, Value: v}
		}

		return nil
	}
}

//////////////////////////////////////////////////////////////
// Greater than
//////////////////////////////////////////////////////////////

var ErrIntegerMultipleOf = errors.New("is not a multiple")

type ErrIntegerMultipleOfDetails[T constraints.Integer] struct {
	Name       string
	MultipleOf T
	Value      T
}

func (e ErrIntegerMultipleOfDetails[T]) Error() string {
	return fmt.Sprintf(
		"%s value of %d is %s of %d",
		e.Name, e.Value, ErrIntegerMultipleOf.Error(), e.MultipleOf,
	)
}

func (e *ErrIntegerMultipleOfDetails[T]) Unwrap() error {
	return ErrIntegerMultipleOf
}

func IntegerMultipleOf[T constraints.Integer](name string, multipleOf T, value T) error {
	if value%multipleOf != 0 {
		return &ErrIntegerMultipleOfDetails[T]{Name: name, MultipleOf: multipleOf, Value: value}
	}

	return nil
}

func IntegerMultipleOfFn[T constraints.Integer](name string, multipleOf T) func(T) error {
	return func(value T) error { return IntegerMultipleOf(name, multipleOf, value) }
}

//////////////////////////////////////////////////////////////
// Zero
//////////////////////////////////////////////////////////////

var ErrIntegerZero = errors.New("is zero")

func IntegerNotZero[T constraints.Integer](name string, value T) error {
	if value == 0 {
		return fmt.Errorf("%s %w", name, ErrIntegerZero)
	}

	return nil
}

func IntegerNotZeroFn[T constraints.Integer](name string) func(T) error {
	return func(value T) error { return IntegerNotZero(name, value) }
}

//////////////////////////////////////////////////////////////
// Not Positive
//////////////////////////////////////////////////////////////

var ErrIntegerPositive = errors.New("is not positive")

type ErrIntegerPositiveDetails[T constraints.Integer] struct {
	Value T
}

func (e ErrIntegerPositiveDetails[T]) Error() string {
	return fmt.Sprintf("%d %s", e.Value, ErrIntegerPositive.Error())
}

func (e *ErrIntegerPositiveDetails[T]) Unwrap() error {
	return ErrIntegerPositive
}

func IntegerPositive[T constraints.Integer](name string, value T) error {
	if value < 0 {
		return &ErrIntegerPositiveDetails[T]{Value: value}
	}

	return nil

}

func IntegerPositiveFn[T constraints.Integer](name string) func(T) error {
	return func(value T) error { return IntegerPositive(name, value) }
}

//////////////////////////////////////////////////////////////
// Not Negative
//////////////////////////////////////////////////////////////

var ErrIntegerNegative = errors.New("is not negative")

type ErrIntegerNegativeDetails[T constraints.Integer] struct {
	Value T
}

func (e ErrIntegerNegativeDetails[T]) Error() string {
	return fmt.Sprintf("%d %s", e.Value, ErrIntegerNegative.Error())
}

func (e *ErrIntegerNegativeDetails[T]) Unwrap() error {
	return ErrIntegerNegative
}

func IntegerNegative[T constraints.Integer](name string, value T) error {
	if value > 0 {
		return &ErrIntegerNegativeDetails[T]{Value: value}
	}

	return nil

}

func IntegerNegativeFn[T constraints.Integer](name string) func(T) error {
	return func(value T) error { return IntegerNegative(name, value) }
}

//////////////////////////////////////////////////////////////
// AnyOf
//////////////////////////////////////////////////////////////

var ErrIntegerAnyOf = errors.New("is not any of")

type ErrIntegerAnyOfDetails[T constraints.Integer] struct {
	Name  string
	OneOf []T
	Value T
}

func (e ErrIntegerAnyOfDetails[T]) Error() string {
	return fmt.Sprintf(
		"%s value of %d %s the following %s",
		e.Name, e.Value, ErrIntegerAnyOf.Error(),
		stringer.Join(", ", stringer.AsStrings(e.OneOf...)...),
	)
}

func (e *ErrIntegerAnyOfDetails[T]) Unwrap() error {
	return ErrIntegerNegative
}

func IntegerAnyOf[T constraints.Integer](name string, value T, elems ...T) error {
	if !slicer.AnyOf(value, elems...) {
		return &ErrIntegerAnyOfDetails[T]{Name: name, Value: value, OneOf: elems}
	}

	return nil
}

func IntegerAnyOfFn[T constraints.Integer](name string, elems ...T) func(T) error {
	return func(value T) error { return IntegerAnyOf(name, value, elems...) }
}

//////////////////////////////////////////////////////////////
// NoneOf
//////////////////////////////////////////////////////////////

var ErrIntegerNoneOf = errors.New("is one of")

type ErrIntegerNoneOfDetails[T constraints.Integer] struct {
	Name   string
	NoneOf []T
	Value  T
}

func (e ErrIntegerNoneOfDetails[T]) Error() string {
	return fmt.Sprintf(
		"%s value of %d %s the following %s",
		e.Name, e.Value, ErrIntegerNoneOf.Error(),
		stringer.Join(", ", stringer.AsStrings(e.NoneOf...)...),
	)
}

func (e *ErrIntegerNoneOfDetails[T]) Unwrap() error {
	return ErrIntegerNegative
}

func IntegerNoneOf[T constraints.Integer](name string, value T, elems ...T) error {
	if !slicer.NoneOf(value, elems...) {
		return &ErrIntegerNoneOfDetails[T]{Name: name, Value: value, NoneOf: elems}
	}

	return nil
}

func IntegerNoneOfFn[T constraints.Integer](name string, elems ...T) func(T) error {
	return func(value T) error { return IntegerNoneOf(name, value, elems...) }
}
