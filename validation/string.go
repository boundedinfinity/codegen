package validation

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// ================================================================================================
// String Empty
// ================================================================================================

var ErrStringEmpty = errors.New("is empty")

func StringNotEmtpy[T ~string](name string, value T) error {
	if value == "" {
		return fmt.Errorf("%s %w", name, ErrStringEmpty)
	}

	return nil
}

func StringNotEmptyFn[T ~string](name string) func(T) error {
	return func(value T) error { return StringNotEmtpy(name, value) }
}

// ================================================================================================
// String Required
// ================================================================================================

var ErrStringRequired = errors.New("is required")

func StringRequired[T ~string](name string, value T) error {
	if value == "" {
		return fmt.Errorf("%s %w", name, ErrStringRequired)
	}

	return nil
}

func StringRequiredFn[T ~string](name string) func(v T) error {
	return func(value T) error { return StringRequired(name, value) }
}

// ================================================================================================
// String Min
// ================================================================================================

var ErrStringLessThanMin = errors.New("length is less than min value")

func StringMin[T ~string](name string, min int, value T) error {
	if len(value) < min {
		return fmt.Errorf("%s value %s %w of %d", name, value, ErrStringLessThanMin, min)
	}

	return nil
}

func StringMinFn[T ~string](name string, min int) func(T) error {
	return func(value T) error { return StringMin(name, min, value) }
}

// ================================================================================================
// String Max
// ================================================================================================

var ErrStringGreaterThanMax = errors.New("length is greater than max value")

func StringMax[T ~string](name string, max int, value T) error {
	if len(value) > max {
		return fmt.Errorf("%s value %s %w of %d", name, value, ErrStringGreaterThanMax, max)
	}

	return nil
}

func StringMaxFn[T ~string](name string, max int) func(T) error {
	return func(value T) error { return StringMin(name, max, value) }
}

// ================================================================================================
// String Regex
// ================================================================================================

var ErrStringDoesNotMatchPattern = errors.New("does not match pattern")

func StringRegex[T ~string](name string, pattern string, value T) error {
	regex := regexp.MustCompile(pattern)

	if !regex.MatchString(string(value)) {
		return fmt.Errorf("%s value %s %w of %s", name, value, ErrStringDoesNotMatchPattern, pattern)
	}

	return nil
}

func StringRegexFn[T ~string](name string, pattern string) func(T) error {
	regex := regexp.MustCompile(pattern)

	return func(value T) error {
		if !regex.MatchString(string(value)) {
			return fmt.Errorf("%s value %s %w of %s", name, value, ErrStringDoesNotMatchPattern, pattern)
		}

		return nil
	}
}

// ================================================================================================
// String UpperCase
// ================================================================================================

var ErrStringNotUpperCase = errors.New("is not upper cased")

func StringUpperCaseFn[T ~string](name string) func(v T) error {
	return func(v T) error {
		if stringer.Capitalize(v) != string(v) {
			return fmt.Errorf("%s value %w", name, ErrStringNotUpperCase)
		}

		return nil
	}
}

// ================================================================================================
// String LowerCase
// ================================================================================================

var ErrStringNotLowerCase = errors.New("is not lower cased")

func StringLowerCase[T ~string](name string) func(v T) error {
	return func(v T) error {
		if stringer.Capitalize(v) != string(v) {
			return fmt.Errorf("%s value %w", name, ErrStringNotLowerCase)
		}

		return nil
	}
}

// ================================================================================================
// String ContainsAny
// ================================================================================================

var ErrStringDoesNotContainAny = errors.New("does not contain given value")

func StringContainsAnyFn[T ~string](name string, elems ...T) func(v T) error {
	return func(v T) error {
		if !stringer.ContainsAny(v, elems...) {
			return fmt.Errorf("%s value %s %w from %s",
				name, v, ErrStringDoesNotContainAny,
				stringer.Join(", ", elems...),
			)
		}

		return nil
	}
}

// ================================================================================================
// String ContainsNone
// ================================================================================================

var ErrStringContainSome = errors.New("does contain given value")

func StringContainsNoneFn[T ~string](name string, elems ...T) func(v T) error {
	return func(v T) error {
		if !stringer.ContainsNone(v, elems...) {
			return fmt.Errorf("%s value %s %w from %s",
				name, v, ErrStringContainSome,
				stringer.Join(", ", elems...),
			)
		}

		return nil
	}
}

// ================================================================================================
// String OneOf
// ================================================================================================

var ErrStringNotOneOf = errors.New("is not one of given value")

func StringOneOfFn[T ~string](name string, elems ...T) func(v T) error {
	return func(v T) error {
		if !stringer.ContainsAny(v, elems...) {
			return fmt.Errorf("%s value %s %w from %s",
				name, v, ErrStringDoesNotContainAny,
				stringer.Join(", ", elems...),
			)
		}

		return nil
	}
}
