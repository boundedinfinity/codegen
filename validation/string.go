package validation

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

var ErrStringRequired = errors.New("is required")

func StringRequired[T ~string](name string) func(v T) error {
	return func(v T) error {
		if v == "" {
			return fmt.Errorf("%s %w", name, ErrStringRequired)
		}

		return nil
	}
}

var ErrStringLessThanMin = errors.New("length is less than min value")

func StringMin[T ~string](name string, n int) func(v T) error {
	return func(v T) error {
		if len(v) < n {
			return fmt.Errorf("%s value %s %w of %d", name, v, ErrStringLessThanMin, n)
		}

		return nil
	}
}

var ErrStringGreaterThanMax = errors.New("length is greater than max value")

func StringMax[T ~string](name string, n int) func(v T) error {
	return func(v T) error {
		if len(v) < n {
			return fmt.Errorf("%s value %s %w of %d", name, v, ErrStringGreaterThanMax, n)
		}

		return nil
	}
}

var ErrStringDoesNotMatchPattern = errors.New("does not match pattern")

func StringRegex[T ~string](name string, pattern string) func(v T) error {
	regex := regexp.MustCompile(pattern)

	return func(v T) error {
		if !regex.MatchString(string(v)) {
			return fmt.Errorf("%s value %s %w of %s", name, v, ErrStringDoesNotMatchPattern, pattern)
		}

		return nil
	}
}

var ErrStringEmpty = errors.New("is empty")

func StringNotEmpty[T ~string](name string) func(v T) error {
	return func(v T) error {
		if v == "" {
			return fmt.Errorf("%s %w", name, ErrStringEmpty)
		}

		return nil
	}
}

var ErrStringNotUpperCase = errors.New("is not upper cased")

func StringUpperCase[T ~string](name string) func(v T) error {
	return func(v T) error {
		if stringer.Capitalize(v) != string(v) {
			return fmt.Errorf("%s value %w", name, ErrStringNotUpperCase)
		}

		return nil
	}
}

var ErrStringNotLowerCase = errors.New("is not lower cased")

func StringLowerCase[T ~string](name string) func(v T) error {
	return func(v T) error {
		if stringer.Capitalize(v) != string(v) {
			return fmt.Errorf("%s value %w", name, ErrStringNotLowerCase)
		}

		return nil
	}
}

var ErrStringDoesNotContainAny = errors.New("does not contain given value")

func StringContainsAny[T ~string](name string, elems ...T) func(v T) error {
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

var ErrStringContainSome = errors.New("does contain given value")

func StringContainsNone[T ~string](name string, elems ...T) func(v T) error {
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

var ErrStringNotOneOf = errors.New("is not one of given value")

func StringOneOf[T ~string](name string, elems ...T) func(v T) error {
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
