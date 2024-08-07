package validation

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// Indexed Name
func in(name string, length, index int) string {
	if length > 0 {
		return fmt.Sprintf("%s[%d]", name, index)
	}

	return name
}

// ================================================================================================
// String Empty
// ================================================================================================

var ErrStringEmpty = errors.New("string is empty")

type ErrStringEmtpyDetails struct {
	Name   string
	Length int
	Index  int
}

func (e *ErrStringEmtpyDetails) Error() string {
	return fmt.Sprintf("%s %s", in(e.Name, e.Length, e.Index), ErrStringEmpty.Error())
}

func (e *ErrStringEmtpyDetails) Unwrap() error {
	return ErrStringEmpty
}

func StringNotEmtpy[T ~string](name string, values ...T) error {
	for i, value := range values {
		if value == "" {
			return &ErrStringEmtpyDetails{Name: name, Index: i, Length: len(values)}
		}
	}

	return nil
}

func StringNotEmptyFn[T ~string](name string) func(...T) error {
	return func(values ...T) error { return StringNotEmtpy(name, values...) }
}

// ================================================================================================
// String Required
// ================================================================================================

var ErrStringRequired = errors.New("string is required")

type ErrStringRequiredDetails[T ~string] struct {
	Name   string
	Length int
	Index  int
	Value  T
}

func (e ErrStringRequiredDetails[T]) Error() string {
	return fmt.Sprintf("%s %s", in(e.Name, e.Length, e.Index), ErrStringRequired.Error())
}

func (e ErrStringRequiredDetails[T]) Unwrap() error {
	return ErrStringRequired
}

func StringRequired[T ~string](name string, values ...T) error {
	for i, value := range values {
		if value == "" {
			return &ErrStringRequiredDetails[T]{
				Name:   name,
				Length: len(values),
				Index:  i,
				Value:  value,
			}
		}
	}

	return nil
}

func StringRequiredFn[T ~string](name string) func(...T) error {
	return func(values ...T) error { return StringRequired(name, values...) }
}

// ================================================================================================
// String Min
// ================================================================================================

var ErrStringMin = errors.New("string length is less than min value")

type ErrStringMinDetails[T ~string] struct {
	Name   string
	Length int
	Index  int
	Value  T
	Min    int
}

func (e ErrStringMinDetails[T]) Error() string {
	return fmt.Sprintf("%s %s", in(e.Name, e.Length, e.Index), ErrStringRequired.Error())
}

func (e ErrStringMinDetails[T]) Unwrap() error {
	return ErrStringMin
}

func StringMin[T ~string](name string, min int, values ...T) error {
	for i, value := range values {
		if len(value) < min {
			return &ErrStringMinDetails[T]{
				Name:   name,
				Length: len(values),
				Index:  i,
				Value:  value,
				Min:    min,
			}
		}
	}

	return nil
}

func StringMinFn[T ~string](name string, min int) func(...T) error {
	return func(values ...T) error { return StringMin(name, min, values...) }
}

// ================================================================================================
// String Max
// ================================================================================================

var ErrStringMax = errors.New("string length is greater than max")

type ErrStringMaxDetails[T ~string] struct {
	Name   string
	Length int
	Index  int
	Value  T
	Max    int
}

func (e ErrStringMaxDetails[T]) Error() string {
	return fmt.Sprintf("%s(%s) %s of %d",
		in(e.Name, e.Length, e.Index), e.Value, ErrStringMax.Error(), e.Max,
	)
}

func (e ErrStringMaxDetails[T]) Unwrap() error {
	return ErrStringMax
}

func StringMax[T ~string](name string, max int, values ...T) error {
	for i, value := range values {
		if len(value) > max {
			return &ErrStringMaxDetails[T]{
				Name:   name,
				Length: len(values),
				Index:  i,
				Value:  value,
				Max:    max,
			}
		}
	}

	return nil
}

func StringMaxFn[T ~string](name string, max int) func(...T) error {
	return func(values ...T) error { return StringMin(name, max, values...) }
}

// ================================================================================================
// String Regex
// ================================================================================================

var ErrStringRegex = errors.New("string does not match pattern")

type ErrStringRegexDetails[T ~string] struct {
	Name    string
	Length  int
	Index   int
	Value   T
	Pattern string
}

func (e ErrStringRegexDetails[T]) Error() string {
	return fmt.Sprintf("%s(%s) %s %s",
		in(e.Name, e.Length, e.Index), e.Value, ErrStringMax.Error(), e.Pattern,
	)
}

func (e ErrStringRegexDetails[T]) Unwrap() error {
	return ErrStringRegex
}

func StringRegex[T ~string](name string, pattern string, values ...T) error {
	regex := regexp.MustCompile(pattern)

	for i, value := range values {
		if !regex.MatchString(string(value)) {
			return &ErrStringRegexDetails[T]{
				Name:    name,
				Pattern: pattern,
				Length:  len(values),
				Index:   i,
				Value:   value,
			}
		}
	}

	return nil
}

func StringRegexFn[T ~string](name string, pattern string) func(...T) error {
	return func(values ...T) error { return StringRegex[T](name, pattern, values...) }
}

// ================================================================================================
// String UpperCase
// ================================================================================================

var ErrStringNotUpperCase = errors.New("string is not upper cased")

type ErrStringNotUpperCaseDetails[T ~string] struct {
	Name   string
	Length int
	Index  int
	Value  T
}

func (e ErrStringNotUpperCaseDetails[T]) Error() string {
	return fmt.Sprintf("%s(%s) %s",
		in(e.Name, e.Length, e.Index), e.Value, ErrStringMax.Error(),
	)
}

func (e ErrStringNotUpperCaseDetails[T]) Unwrap() error {
	return ErrStringRegex
}

func StringUpperCase[T ~string](name string, values ...T) error {
	for i, value := range values {
		if stringer.Capitalize(value) != string(value) {
			return &ErrStringNotUpperCaseDetails[T]{
				Name:   name,
				Length: len(values),
				Index:  i,
				Value:  value,
			}
		}
	}

	return nil

}

func StringUpperCaseFn[T ~string](name string) func(values ...T) error {
	return func(values ...T) error { return StringUpperCase(name, values...) }
}

// ================================================================================================
// String LowerCase
// ================================================================================================

var ErrStringNotLowerCase = errors.New("is not lower cased")

func StringLowerCaseFn[T ~string](name string) func(...T) error {
	return func(values ...T) error {
		for _, value := range values {
			if stringer.Capitalize(value) != string(value) {
				return fmt.Errorf("%s value %w", name, ErrStringNotLowerCase)
			}
		}

		return nil
	}
}

// ================================================================================================
// String ContainsAny
// ================================================================================================

var ErrStringDoesNotContainAny = errors.New("does not contain given value")

func StringContainsAnyFn[T ~string](name string, elems ...T) func(...T) error {
	return func(values ...T) error {
		for _, value := range values {
			if !stringer.ContainsAny(value, elems...) {
				return fmt.Errorf("%s value %s %w from %s",
					name, value, ErrStringDoesNotContainAny,
					stringer.Join(", ", elems...),
				)
			}
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
