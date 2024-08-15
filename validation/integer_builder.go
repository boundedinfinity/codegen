package validation

import (
	"boundedinfinity/codegen/model"
	"errors"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

func Integer[T ~int](name string) *integerValidations[T] {
	return &integerValidations[T]{
		name:        name,
		validations: []func(v T) error{},
	}
}

type integerValidations[T ~int] struct {
	name        string
	validations []func(v T) error
}

func (t integerValidations[T]) Validate(v T) error {
	errs := []error{}

	for _, validation := range t.validations {
		if err := validation(v); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (t *integerValidations[T]) Min(n T) *integerValidations[T] {
	t.validations = append(t.validations, IntegerMinFn(t.name, n))
	return t
}

func (t *integerValidations[T]) Max(n T) *integerValidations[T] {
	t.validations = append(t.validations, IntegerMaxFn(t.name, n))
	return t
}

func (t *integerValidations[T]) MinMax(min, max T) *integerValidations[T] {
	t.Range(model.Range[T]{Min: optioner.Some(min), Max: optioner.Some(max)})
	return t
}

func (t *integerValidations[T]) Range(rng model.Range[T]) *integerValidations[T] {
	t.validations = append(t.validations, IntegerRange(t.name, rng))
	return t
}

func (t *integerValidations[T]) MultipleOf(n T) *integerValidations[T] {
	t.validations = append(t.validations, IntegerMultipleOfFn(t.name, n))
	return t
}

func (t *integerValidations[T]) NotZero() *integerValidations[T] {
	t.validations = append(t.validations, IntegerNotZeroFn[T](t.name))
	return t
}

func (t *integerValidations[T]) Positive() *integerValidations[T] {
	t.validations = append(t.validations, IntegerPositiveFn[T](t.name))
	return t
}

func (t *integerValidations[T]) Negative() *integerValidations[T] {
	t.validations = append(t.validations, IntegerNegativeFn[T](t.name))
	return t
}
