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

func StringNotEmtpy[T ~string](name string, value T) error {
	if value == "" {
		return &ErrStringEmtpyDetails{Name: name}
	}

	return nil
}

func StringNotEmptyFn[T ~string](name string) func(T) error {
	return func(value T) error { return StringNotEmtpy(name, value) }
}

func StringsNotEmtpy[T ~string](name string, values ...T) error {
	for i, value := range values {
		if value == "" {
			return &ErrStringEmtpyDetails{Name: name, Index: i, Length: len(values)}
		}
	}

	return nil
}

func StringsNotEmptyFn[T ~string](name string) func(...T) error {
	return func(values ...T) error { return StringsNotEmtpy(name, values...) }
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

func StringRequired[T ~string](name string, value T) error {

	if value == "" {
		return &ErrStringRequiredDetails[T]{
			Name:  name,
			Value: value,
		}
	}

	return nil
}

func StringRequiredFn[T ~string](name string) func(T) error {
	return func(value T) error { return StringRequired(name, value) }
}

func StringsRequired[T ~string](name string, values ...T) error {
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

func StringsRequiredFn[T ~string](name string) func(...T) error {
	return func(values ...T) error { return StringsRequired(name, values...) }
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

func StringMin[T ~string](name string, min int, value T) error {
	if len(value) < min {
		return &ErrStringMinDetails[T]{
			Name:   name,
			Length: len(value),
			Value:  value,
			Min:    min,
		}
	}

	return nil
}

func StringMinFn[T ~string](name string, min int) func(T) error {
	return func(value T) error { return StringMin(name, min, value) }
}

func StringsMin[T ~string](name string, min int, values ...T) error {
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

func StringsMinFn[T ~string](name string, min int) func(...T) error {
	return func(values ...T) error { return StringsMin(name, min, values...) }
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

func StringMax[T ~string](name string, max int, value T) error {
	if len(value) > max {
		return &ErrStringMaxDetails[T]{
			Name:   name,
			Length: len(value),
			Value:  value,
			Max:    max,
		}
	}

	return nil
}

func StringMaxFn[T ~string](name string, max int) func(T) error {
	return func(value T) error { return StringMax(name, max, value) }
}

func StringsMax[T ~string](name string, max int, values ...T) error {
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

func StringsMaxFn[T ~string](name string, max int) func(...T) error {
	return func(values ...T) error { return StringsMax(name, max, values...) }
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

func StringRegex[T ~string](name string, pattern string, value T) error {
	regex := regexp.MustCompile(pattern)

	if !regex.MatchString(string(value)) {
		return &ErrStringRegexDetails[T]{
			Name:    name,
			Pattern: pattern,
			Length:  len(value),
			Value:   value,
		}
	}

	return nil
}

func StringRegexFn[T ~string](name string, pattern string) func(T) error {
	return func(value T) error { return StringRegex[T](name, pattern, value) }
}

func StringsRegex[T ~string](name string, pattern string, values ...T) error {
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

func StringsRegexFn[T ~string](name string, pattern string) func(...T) error {
	return func(values ...T) error { return StringsRegex[T](name, pattern, values...) }
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
	return ErrStringNotUpperCase
}

func StringUpperCase[T ~string](name string, value T) error {
	if stringer.Uppercase(value) != string(value) {
		return &ErrStringNotUpperCaseDetails[T]{
			Name:   name,
			Length: len(value),
			Value:  value,
		}
	}

	return nil

}

func StringUpperCaseFn[T ~string](name string) func(value T) error {
	return func(value T) error { return StringUpperCase(name, value) }
}

func StringsUpperCase[T ~string](name string, values ...T) error {
	for i, value := range values {
		if stringer.Uppercase(value) != string(value) {
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

func StringsUpperCaseFn[T ~string](name string) func(values ...T) error {
	return func(values ...T) error { return StringsUpperCase(name, values...) }
}

// ================================================================================================
// String LowerCase
// ================================================================================================

var ErrStringNotLowerCase = errors.New("is not lower cased")

type ErrStringNotLowerCaseDetails[T ~string] struct {
	Name   string
	Length int
	Index  int
	Value  T
}

func (e ErrStringNotLowerCaseDetails[T]) Error() string {
	return fmt.Sprintf("%s(%s) %s",
		in(e.Name, e.Length, e.Index), e.Value, ErrStringMax.Error(),
	)
}

func (e ErrStringNotLowerCaseDetails[T]) Unwrap() error {
	return ErrStringNotLowerCase
}

func StringLowerCase[T ~string](name string, value T) error {

	if stringer.Lowercase(value) != string(value) {
		return &ErrStringNotLowerCaseDetails[T]{
			Name:  name,
			Value: value,
		}
	}

	return nil
}

func StringLowerCaseFn[T ~string](name string) func(T) error {
	return func(value T) error { return StringLowerCase(name, value) }
}

func StringsLowerCase[T ~string](name string, values ...T) error {
	for i, value := range values {
		if stringer.Lowercase(value) != string(value) {
			return &ErrStringNotLowerCaseDetails[T]{
				Name:   name,
				Length: len(values),
				Index:  i,
				Value:  value,
			}
		}
	}

	return nil
}

func StringsLowerCaseFn[T ~string](name string) func(...T) error {
	return func(values ...T) error { return StringsLowerCase(name, values...) }
}

// ================================================================================================
// String ContainsAny
// ================================================================================================

var ErrStringContainsAny = errors.New("string does any of the required values")

type ErrStringContainsAnyDetails[T ~string, U ~string] struct {
	Name   string
	Length int
	Index  int
	AnyOf  []T
	Value  U
}

func (e ErrStringContainsAnyDetails[T, U]) Error() string {
	return fmt.Sprintf("%s(%s) %s from %s",
		in(e.Name, e.Length, e.Index), e.Value,
		ErrStringMax.Error(), stringer.Join(", ", e.AnyOf...),
	)
}

func (e ErrStringContainsAnyDetails[T, U]) Unwrap() error {
	return ErrStringContainsAny
}

func StringContainsAny[T ~string, U ~string](name string, elems []T, values ...U) error {
	for i, value := range values {
		if !stringer.ContainsAny(value, elems...) {
			return &ErrStringContainsAnyDetails[T, U]{
				Name:   name,
				Length: len(values),
				Index:  i,
				AnyOf:  elems,
				Value:  value,
			}
		}
	}

	return nil
}

func StringContainsAnyFn[T ~string, U ~string](name string, elems ...T) func(...U) error {
	return func(values ...U) error { return StringContainsAny(name, elems, values...) }
}

// ================================================================================================
// String ContainsNone
// ================================================================================================

var ErrStringContainsNone = errors.New("string contains at least one of invalid values")

type ErrStringContainsNoneDetails[T ~string, U ~string] struct {
	Name   string
	Length int
	Index  int
	NoneOf []T
	Value  U
}

func (e ErrStringContainsNoneDetails[T, U]) Error() string {
	return fmt.Sprintf("%s(%s) %s from %s",
		in(e.Name, e.Length, e.Index), e.Value,
		ErrStringMax.Error(), stringer.Join(", ", e.NoneOf...),
	)
}

func (e ErrStringContainsNoneDetails[T, U]) Unwrap() error {
	return ErrStringContainsNone
}

func StringContainsNone[T ~string, U ~string](name string, elems []T, values ...U) error {
	for i, value := range values {
		if !stringer.ContainsNone(value, elems...) {
			return &ErrStringContainsNoneDetails[T, U]{
				Name:   name,
				Length: len(values),
				Index:  i,
				NoneOf: elems,
				Value:  value,
			}
		}
	}

	return nil

}

func StringContainsNoneFn[T ~string, U ~string](name string, elems ...T) func(...U) error {
	return func(values ...U) error { return StringContainsNone(name, elems, values...) }
}

// ================================================================================================
// String OneOf
// ================================================================================================

var ErrStringNotOneOf = errors.New("is not one of given value")

func StringOneOfFn[T ~string](name string, elems ...T) func(v T) error {
	return func(v T) error {
		if !stringer.ContainsAny(v, elems...) {
			return fmt.Errorf("%s value %s %w from %s",
				name, v, ErrStringNotOneOf,
				stringer.Join(", ", elems...),
			)
		}

		return nil
	}
}
