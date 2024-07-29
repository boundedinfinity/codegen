package validation

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func String[T ~string](name string) *stringValidations[T] {
	return &stringValidations[T]{
		name:        name,
		validations: []func(v T) error{},
	}
}

type stringValidations[T ~string] struct {
	name        string
	validations []func(v T) error
}

func (t stringValidations[T]) Validate(v T) error {
	errs := []error{}

	for _, validation := range t.validations {
		if err := validation(v); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (t *stringValidations[T]) Min(n int) *stringValidations[T] {
	t.validations = append(t.validations, StringMin[T](t.name, n))
	return t
}

func (t *stringValidations[T]) Max(n int) *stringValidations[T] {
	t.validations = append(t.validations, StringMax[T](t.name, n))
	return t
}

func (t *stringValidations[T]) Regex(pattern string) *stringValidations[T] {
	t.validations = append(t.validations, StringRegex[T](t.name, pattern))
	return t
}

func (t *stringValidations[T]) NotEmpty(pattern string) *stringValidations[T] {
	t.validations = append(t.validations, StringNotEmpty[T](t.name))
	return t
}

func (t *stringValidations[T]) UpperCase(pattern string) *stringValidations[T] {
	t.validations = append(t.validations, StringUpperCase[T](t.name))
	return t
}

func (t *stringValidations[T]) LowerCase(pattern string) *stringValidations[T] {
	t.validations = append(t.validations, StringLowerCase[T](t.name))
	return t
}

func (t *stringValidations[T]) Required() *stringValidations[T] {
	t.validations = append(t.validations, StringRequired[T](t.name))
	return t
}

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
