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

func (t floatValidations[T]) Validate(v T) error {
	errs := []error{}

	for _, validation := range t.validations {
		if err := validation(v); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (t *floatValidations[T]) Min(n T) *floatValidations[T] {
	t.validations = append(t.validations, FloatMin[T](t.name, n))
	return t
}

func (t *floatValidations[T]) Max(n T) *floatValidations[T] {
	t.validations = append(t.validations, FloatMax[T](t.name, n))
	return t
}

// func (t *floatValidations[T]) MultipleOf(n T) *floatValidations[T] {
// 	t.validations = append(t.validations, FloatMultipleOf[T](t.name, n))
// 	return t
// }

func (t *floatValidations[T]) NotZero() *floatValidations[T] {
	t.validations = append(t.validations, FloatNotZero[T](t.name))
	return t
}

func (t *floatValidations[T]) Positive() *floatValidations[T] {
	t.validations = append(t.validations, FloatPositive[T](t.name))
	return t
}

func (t *floatValidations[T]) Negative() *floatValidations[T] {
	t.validations = append(t.validations, FloatNegative[T](t.name))
	return t
}
