package validation

import "errors"

func Float[T ~float32 | ~float64](name string) *floatValidations[T] {
	return &floatValidations[T]{
		name:        name,
		validations: []func(v T) error{},
	}
}

type floatValidations[T ~float32 | ~float64] struct {
	name        string
	validations []func(v T) error
}

func (this floatValidations[T]) Validate(v T) error {
	errs := []error{}

	for _, validation := range this.validations {
		if err := validation(v); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (this *floatValidations[T]) Min(n T) *floatValidations[T] {
	this.validations = append(this.validations, FloatMin[T](this.name, n))
	return this
}

func (this *floatValidations[T]) Max(n T) *floatValidations[T] {
	this.validations = append(this.validations, FloatMax[T](this.name, n))
	return this
}

// func (this *floatValidations[T]) MultipleOf(n T) *floatValidations[T] {
// 	this.validations = append(this.validations, FloatMultipleOf[T](this.name, n))
// 	return this
// }

func (this *floatValidations[T]) NotZero() *floatValidations[T] {
	this.validations = append(this.validations, FloatNotZero[T](this.name))
	return this
}

func (this *floatValidations[T]) Positive() *floatValidations[T] {
	this.validations = append(this.validations, FloatPositive[T](this.name))
	return this
}

func (this *floatValidations[T]) Negative() *floatValidations[T] {
	this.validations = append(this.validations, FloatNegative[T](this.name))
	return this
}
