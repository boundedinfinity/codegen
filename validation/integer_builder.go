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

func (this integerValidations[T]) Validate(v T) error {
	errs := []error{}

	for _, validation := range this.validations {
		if err := validation(v); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (this *integerValidations[T]) Min(n T) *integerValidations[T] {
	this.validations = append(this.validations, IntegerMinFn(this.name, n))
	return this
}

func (this *integerValidations[T]) Max(n T) *integerValidations[T] {
	this.validations = append(this.validations, IntegerMaxFn(this.name, n))
	return this
}

func (this *integerValidations[T]) MinMax(min, max T) *integerValidations[T] {
	this.Range(model.Range[T]{Min: optioner.Some(min), Max: optioner.Some(max)})
	return this
}

func (this *integerValidations[T]) Range(rng model.Range[T]) *integerValidations[T] {
	this.validations = append(this.validations, IntegerRange(this.name, rng))
	return this
}

func (this *integerValidations[T]) MultipleOf(n T) *integerValidations[T] {
	this.validations = append(this.validations, IntegerMultipleOfFn(this.name, n))
	return this
}

func (this *integerValidations[T]) NotZero() *integerValidations[T] {
	this.validations = append(this.validations, IntegerNotZeroFn[T](this.name))
	return this
}

func (this *integerValidations[T]) Positive() *integerValidations[T] {
	this.validations = append(this.validations, IntegerPositiveFn[T](this.name))
	return this
}

func (this *integerValidations[T]) Negative() *integerValidations[T] {
	this.validations = append(this.validations, IntegerNegativeFn[T](this.name))
	return this
}
