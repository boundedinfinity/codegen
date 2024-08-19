package validation

import (
	"boundedinfinity/codegen/model"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"golang.org/x/exp/constraints"
)

type integerValidations[T constraints.Integer] struct {
	validator *validator[T]
}

var _ Validater[int] = &integerValidations[int]{}

func Integer[T constraints.Integer](name string) *integerValidations[T] {
	return &integerValidations[T]{validator: &validator[T]{name: name}}
}

func (this *integerValidations[T]) Validate(value T) error {
	return this.validator.Validate(value)
}

func (this *integerValidations[T]) Min(n T) *integerValidations[T] {
	this.validator.append(IntegerMinFn[T](this.validator.name, n))
	return this
}

func (this *integerValidations[T]) Max(n T) *integerValidations[T] {
	this.validator.append(IntegerMaxFn[T](this.validator.name, n))
	return this
}

func (this *integerValidations[T]) MinMax(min, max T) *integerValidations[T] {
	this.Range(model.Range[T]{Min: optioner.Some(min), Max: optioner.Some(max)})
	return this
}

func (this *integerValidations[T]) Range(rng model.Range[T]) *integerValidations[T] {
	this.validator.append(IntegerRange[T](this.validator.name, rng))
	return this
}

func (this *integerValidations[T]) MultipleOf(n T) *integerValidations[T] {
	this.validator.append(IntegerMultipleOfFn[T](this.validator.name, n))
	return this
}

func (this *integerValidations[T]) NotZero() *integerValidations[T] {
	this.validator.append(IntegerNotZeroFn[T](this.validator.name))
	return this
}

func (this *integerValidations[T]) Positive() *integerValidations[T] {
	this.validator.append(IntegerPositiveFn[T](this.validator.name))
	return this
}

func (this *integerValidations[T]) Negative() *integerValidations[T] {
	this.validator.append(IntegerNegativeFn[T](this.validator.name))
	return this
}
