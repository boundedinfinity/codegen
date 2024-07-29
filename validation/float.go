package validation

import (
	"errors"
	"fmt"
)

var ErrFloatLessThanMin = errors.New("is less than min value")

func FloatMin[T ~float32 | ~float64](name string, n T) func(v T) error {
	return func(v T) error {
		if v < n {
			return fmt.Errorf("%s value %f %w of %f", name, v, ErrFloatLessThanMin, n)
		}

		return nil
	}
}

var ErrFloatGreaterThanMax = errors.New("is greater than max value")

func FloatMax[T ~float32 | ~float64](name string, n T) func(v T) error {
	return func(v T) error {
		if v < n {
			return fmt.Errorf("%s value %f %w of %f", name, v, ErrFloatGreaterThanMax, n)
		}

		return nil
	}
}

// var ErrFloatNotMultipleOf = errors.New("is not a multiple")

// func FloatMultipleOf[T ~float32 | ~float64](name string, n T) func(v T) error {
// 	return func(v T) error {
// 		if v%n != 0 {
// 			return fmt.Errorf("%s value %f %w of %f", name, v, ErrStringDoesNotMatchPattern, n)
// 		}

// 		return nil
// 	}
// }

var ErrFloatZero = errors.New("is zero")

func FloatNotZero[T ~float32 | ~float64](name string) func(v T) error {
	return func(v T) error {
		if v == 0 {
			return fmt.Errorf("%s %w", name, ErrFloatZero)
		}

		return nil
	}
}

var ErrFloatNotPositive = errors.New("is not positive")

func FloatPositive[T ~float32 | ~float64](name string) func(v T) error {
	return func(v T) error {
		if v < 0 {
			return fmt.Errorf("%s value v %w", name, ErrFloatNotPositive)
		}

		return nil
	}
}

var ErrFloatNotNegative = errors.New("is not negative")

func FloatNegative[T ~float32 | ~float64](name string) func(v T) error {
	return func(v T) error {
		if v > 0 {
			return fmt.Errorf("%s value v %w", name, ErrFloatNotNegative)
		}

		return nil
	}
}
