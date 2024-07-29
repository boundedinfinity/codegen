package validation

import "errors"

var ValidationError = errors.New("validation error")

func newValidations[T any](name string) *validations[T] {
	return &validations[T]{
		name:        name,
		validations: []func(v T) error{},
	}
}

type validations[T any] struct {
	name        string
	validations []func(v T) error
}

func (t validations[T]) Validate(v T) error {
	errs := []error{}

	for _, validation := range t.validations {
		if err := validation(v); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}
