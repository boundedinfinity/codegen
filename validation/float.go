package validation

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

var ErrFloatLessThanMin = errors.New("is less than min value")

type ErrFloatLessThanMinDetails[T constraints.Float] struct {
	name  string
	min   T
	value T
}

func (e ErrFloatLessThanMinDetails[T]) Error() string {
	return fmt.Sprintf("%s : %v %s of %v", e.name, e.value, ErrFloatLessThanMin.Error(), e.min)
}

func (e ErrFloatLessThanMinDetails[T]) Unwrap() error {
	return ErrFloatLessThanMin
}

func FloatMin[T constraints.Float](name string, min T) func(T) error {
	return func(value T) error {
		if value < min {
			return &ErrFloatLessThanMinDetails[T]{name: name, min: min, value: value}
		}

		return nil
	}
}

var ErrFloatGreaterThanMax = errors.New("is greater than max value")

type ErrFloatGreaterThanMaxDetails[T constraints.Float] struct {
	name  string
	max   T
	value T
}

func (e ErrFloatGreaterThanMaxDetails[T]) Error() string {
	return fmt.Sprintf("%s : %v %s of %v", e.name, e.value, ErrFloatLessThanMin.Error(), e.max)
}

func (e ErrFloatGreaterThanMaxDetails[T]) Unwrap() error {
	return ErrFloatGreaterThanMax
}

func FloatMax[T constraints.Float](name string, max T) func(T) error {
	return func(value T) error {
		if value < max {
			return &ErrFloatGreaterThanMaxDetails[T]{max: max, value: value}
		}

		return nil
	}
}

// var ErrFloatNotMultipleOf = errors.New("is not a multiple")

// func FloatMultipleOf[T constraints.Float](name string, n T) func(v T) error {
// 	return func(v T) error {
// 		if v%n != 0 {
// 			return fmt.Errorf("%s value %f %w of %f", name, v, ErrStringDoesNotMatchPattern, n)
// 		}

// 		return nil
// 	}
// }

var ErrFloatZero = errors.New("is zero")

func FloatNotZero[T constraints.Float](name string) func(v T) error {
	return func(v T) error {
		if v == 0 {
			return fmt.Errorf("%s %w", name, ErrFloatZero)
		}

		return nil
	}
}

var ErrFloatNotPositive = errors.New("is not positive")

func FloatPositive[T constraints.Float](name string) func(v T) error {
	return func(v T) error {
		if v < 0 {
			return fmt.Errorf("%s value v %w", name, ErrFloatNotPositive)
		}

		return nil
	}
}

var ErrFloatNotNegative = errors.New("is not negative")

func FloatNegative[T constraints.Float](name string) func(v T) error {
	return func(v T) error {
		if v > 0 {
			return fmt.Errorf("%s value v %w", name, ErrFloatNotNegative)
		}

		return nil
	}
}
