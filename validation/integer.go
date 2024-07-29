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

func IntegerMin[T ~int](name string, min T) func(v T) error {
	return func(v T) error {
		if v < min {
			return &errIntegerLessThanMin[T]{name, min}
		}

		return nil
	}
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

func IntegerMax[T ~int](name string, n T) func(v T) error {
	return func(v T) error {
		if v < n {
			return &errIntegerGreaterThanMax[T]{Name: name, Max: n}
		}

		return nil
	}
}

//////////////////////////////////////////////////////////////
// Greater than
//////////////////////////////////////////////////////////////

var ErrIntegerOutOfRange = errors.New("out of range")

type errIntegerNotInRange[T ~int] struct {
	Name  string
	Range model.NumberRange[T]
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

func IntegerRange[T ~int](name string, rng model.NumberRange[T]) func(v T) error {
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

func IntegerMultipleOf[T ~int](name string, multipleOf T) func(v T) error {
	return func(v T) error {
		if v%multipleOf != 0 {
			return &errIntegerNotMultipleOf[T]{Name: name, MultipleOf: multipleOf, Value: v}
		}

		return nil
	}
}

//////////////////////////////////////////////////////////////
// Zero
//////////////////////////////////////////////////////////////

var ErrIntegerZero = errors.New("is zero")

func IntegerNotZero[T ~int](name string) func(v T) error {
	return func(v T) error {
		if v == 0 {
			return fmt.Errorf("%s %w", name, ErrIntegerZero)
		}

		return nil
	}
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

func IntegerPositive[T ~int](name string) func(v T) error {
	return func(v T) error {
		if v < 0 {
			return &errIntegerNotPositive[T]{Value: v}
		}

		return nil
	}
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

func IntegerNegative[T ~int](name string) func(v T) error {
	return func(v T) error {
		if v > 0 {
			return &errIntegerNotNegative[T]{Value: v}
		}

		return nil
	}
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

func IntegerAnyOf[T ~int](name string, elems ...T) func(v T) error {
	return func(v T) error {
		if !slicer.AnyOf(v, elems...) {
			return &errIntegerNotAnyOf[T]{Name: name, Value: v, OneOf: elems}
		}

		return nil
	}
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

func IntegerNoneOf[T ~int](name string, elems ...T) func(v T) error {
	return func(v T) error {
		if !slicer.NoneOf(v, elems...) {
			return &errIntegerNotNoneOf[T]{Name: name, Value: v, OneOf: elems}
		}

		return nil
	}
}
