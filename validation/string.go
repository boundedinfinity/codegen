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
			return fmt.Errorf("%s value v %w", name, ErrStringNotUpperCase)
		}

		return nil
	}
}

var ErrStringNotLowerCase = errors.New("is not lower cased")

func StringLowerCase[T ~string](name string) func(v T) error {
	return func(v T) error {
		if stringer.Capitalize(v) != string(v) {
			return fmt.Errorf("%s value v %w", name, ErrStringNotLowerCase)
		}

		return nil
	}
}
