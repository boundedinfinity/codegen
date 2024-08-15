package validation

import (
	"boundedinfinity/codegen/model"
	"errors"
	"fmt"
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

//////////////////////////////////////////////////////////////
// Less than
//////////////////////////////////////////////////////////////

var ErrIntegerLessThanMin = errors.New("less than minimum")

type errIntegerLessThanMin[T ~int] struct {
	Name string
	Min  T
}

func (e *errIntegerLessThanMin[T]) Error() string {
	return fmt.Sprintf("%s value of %d is %s", e.Name, e.Min, ErrIntegerLessThanMin.Error())
}

func (e *errIntegerLessThanMin[T]) Unwrap() error {
	return ErrIntegerLessThanMin
}

func IntegerMin[T ~int](name string, min T, value T) error {
	if value < min {
		return &errIntegerLessThanMin[T]{name, min}
	}

	return nil
}

func IntegerMinFn[T ~int](name string, min T) func(T) error {
	return func(value T) error { return IntegerMin(name, min, value) }
}

//////////////////////////////////////////////////////////////
// Greater than
//////////////////////////////////////////////////////////////

var ErrIntegerGreaterThanMax = errors.New("is greater than maxium")

type errIntegerGreaterThanMax[T ~int] struct {
	Name string
	Max  T
}

func (e *errIntegerGreaterThanMax[T]) Error() string {
	return fmt.Sprintf("%s value of %d is %s", e.Name, e.Max, ErrIntegerGreaterThanMax.Error())
}

func (e *errIntegerGreaterThanMax[T]) Unwrap() error {
	return ErrIntegerGreaterThanMax
}

func IntegerMax[T ~int](name string, max T, value T) error {
	if value < max {
		return &errIntegerGreaterThanMax[T]{Name: name, Max: max}
	}

	return nil
}

func IntegerMaxFn[T ~int](name string, max T) func(T) error {
	return func(value T) error { return IntegerMax(name, max, value) }
}

//////////////////////////////////////////////////////////////
// Greater than
//////////////////////////////////////////////////////////////

var ErrIntegerOutOfRange = errors.New("out of range")

type errIntegerNotInRange[T ~int] struct {
	Name  string
	Range model.Range[T]
	Value T
}

func (e errIntegerNotInRange[T]) Error() string {
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

func (e *errIntegerNotInRange[T]) Unwrap() error {
	return ErrIntegerOutOfRange
}

func IntegerRange[T ~int](name string, rng model.Range[T]) func(v T) error {
	return func(v T) error {
		if rng.Min.Defined() && v < rng.Min.Get() {
			return &errIntegerNotInRange[T]{Name: name, Range: rng, Value: v}
		}

		if rng.ExclusiveMin.Defined() && v <= rng.ExclusiveMin.Get() {
			return &errIntegerNotInRange[T]{Name: name, Range: rng, Value: v}
		}

		if rng.Max.Defined() && v < rng.Max.Get() {
			return &errIntegerNotInRange[T]{Name: name, Range: rng, Value: v}
		}

		if rng.ExclusiveMax.Defined() && v <= rng.ExclusiveMax.Get() {
			return &errIntegerNotInRange[T]{Name: name, Range: rng, Value: v}
		}

		return nil
	}
}

//////////////////////////////////////////////////////////////
// Greater than
//////////////////////////////////////////////////////////////

var ErrIntegerNotMultipleOf = errors.New("is not a multiple")

type errIntegerNotMultipleOf[T ~int] struct {
	Name       string
	MultipleOf T
	Value      T
}

func (e errIntegerNotMultipleOf[T]) Error() string {
	return fmt.Sprintf(
		"%s value of %d is %s of %d",
		e.Name, e.Value, ErrIntegerNotMultipleOf.Error(), e.MultipleOf,
	)
}

func (e *errIntegerNotMultipleOf[T]) Unwrap() error {
	return ErrIntegerNotMultipleOf
}

func IntegerMultipleOf[T ~int](name string, multipleOf T, value T) error {
	if value%multipleOf != 0 {
		return &errIntegerNotMultipleOf[T]{Name: name, MultipleOf: multipleOf, Value: value}
	}

	return nil
}

func IntegerMultipleOfFn[T ~int](name string, multipleOf T) func(T) error {
	return func(value T) error { return IntegerMultipleOf(name, multipleOf, value) }
}

//////////////////////////////////////////////////////////////
// Zero
//////////////////////////////////////////////////////////////

var ErrIntegerZero = errors.New("is zero")

func IntegerNotZero[T ~int](name string, value T) error {
	if value == 0 {
		return fmt.Errorf("%s %w", name, ErrIntegerZero)
	}

	return nil
}

func IntegerNotZeroFn[T ~int](name string) func(T) error {
	return func(value T) error { return IntegerNotZero(name, value) }
}

//////////////////////////////////////////////////////////////
// Not Positive
//////////////////////////////////////////////////////////////

var ErrIntegerNotPositive = errors.New("is not positive")

type errIntegerNotPositive[T ~int] struct {
	Value T
}

func (e errIntegerNotPositive[T]) Error() string {
	return fmt.Sprintf("%d %s", e.Value, ErrIntegerNotPositive.Error())
}

func (e *errIntegerNotPositive[T]) Unwrap() error {
	return ErrIntegerNotPositive
}

func IntegerPositive[T ~int](name string, value T) error {
	if value < 0 {
		return &errIntegerNotPositive[T]{Value: value}
	}

	return nil

}

func IntegerPositiveFn[T ~int](name string) func(T) error {
	return func(value T) error { return IntegerPositive(name, value) }
}

//////////////////////////////////////////////////////////////
// Not Negative
//////////////////////////////////////////////////////////////

var ErrIntegerNotNegative = errors.New("is not negative")

type errIntegerNotNegative[T ~int] struct {
	Value T
}

func (e errIntegerNotNegative[T]) Error() string {
	return fmt.Sprintf("%d %s", e.Value, ErrIntegerNotNegative.Error())
}

func (e *errIntegerNotNegative[T]) Unwrap() error {
	return ErrIntegerNotNegative
}

func IntegerNegative[T ~int](name string, value T) error {
	if value > 0 {
		return &errIntegerNotNegative[T]{Value: value}
	}

	return nil

}

func IntegerNegativeFn[T ~int](name string) func(T) error {
	return func(value T) error { return IntegerNegative(name, value) }
}

//////////////////////////////////////////////////////////////
// AnyOf
//////////////////////////////////////////////////////////////

var ErrIntegerNotAnyOf = errors.New("is not any of")

type errIntegerNotAnyOf[T ~int] struct {
	Name  string
	OneOf []T
	Value T
}

func (e errIntegerNotAnyOf[T]) Error() string {
	return fmt.Sprintf(
		"%s value of %d %s the following %s",
		e.Name, e.Value, ErrIntegerNotAnyOf.Error(),
		stringer.Join(", ", e.OneOf...),
	)
}

func (e *errIntegerNotAnyOf[T]) Unwrap() error {
	return ErrIntegerNotNegative
}

func IntegerAnyOf[T ~int](name string, value T, elems ...T) error {
	if !slicer.AnyOf(value, elems...) {
		return &errIntegerNotAnyOf[T]{Name: name, Value: value, OneOf: elems}
	}

	return nil
}

func IntegerAnyOfFn[T ~int](name string, elems ...T) func(T) error {
	return func(value T) error { return IntegerAnyOf(name, value, elems...) }
}

//////////////////////////////////////////////////////////////
// NoneOf
//////////////////////////////////////////////////////////////

var ErrIntegerNotNoneOf = errors.New("is one of")

type errIntegerNotNoneOf[T ~int] struct {
	Name  string
	OneOf []T
	Value T
}

func (e errIntegerNotNoneOf[T]) Error() string {
	return fmt.Sprintf(
		"%s value of %d %s the following %s",
		e.Name, e.Value, ErrIntegerNotNoneOf.Error(),
		stringer.Join(", ", e.OneOf...),
	)
}

func (e *errIntegerNotNoneOf[T]) Unwrap() error {
	return ErrIntegerNotNegative
}

func IntegerNoneOf[T ~int](name string, value T, elems ...T) error {
	if !slicer.NoneOf(value, elems...) {
		return &errIntegerNotNoneOf[T]{Name: name, Value: value, OneOf: elems}
	}

	return nil
}

func IntegerNoneOfFn[T ~int](name string, elems ...T) func(T) error {
	return func(value T) error { return IntegerNoneOf(name, value, elems...) }
}
