package validation

import "errors"

func String[T ~string](name string) *stringValidations[T] {
	return &stringValidations[T]{
		name:        name,
		validations: []func(...T) error{},
	}
}

type stringValidations[T ~string] struct {
	name        string
	validations []func(...T) error
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
	t.validations = append(t.validations, StringMinFn[T](t.name, n))
	return t
}

func (t *stringValidations[T]) Max(n int) *stringValidations[T] {
	t.validations = append(t.validations, StringMaxFn[T](t.name, n))
	return t
}

func (t *stringValidations[T]) Regex(pattern string) *stringValidations[T] {
	t.validations = append(t.validations, StringRegexFn[T](t.name, pattern))
	return t
}

func (t *stringValidations[T]) NotEmpty(pattern string) *stringValidations[T] {
	t.validations = append(t.validations, StringNotEmptyFn[T](t.name))
	return t
}

func (t *stringValidations[T]) UpperCase(pattern string) *stringValidations[T] {
	t.validations = append(t.validations, StringUpperCaseFn[T](t.name))
	return t
}

func (t *stringValidations[T]) LowerCase(pattern string) *stringValidations[T] {
	t.validations = append(t.validations, StringLowerCaseFn[T](t.name))
	return t
}

func (t *stringValidations[T]) Required() *stringValidations[T] {
	t.validations = append(t.validations, StringRequiredFn[T](t.name))
	return t
}
