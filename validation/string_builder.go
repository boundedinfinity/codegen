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

func (this stringValidations[T]) Validate(v T) error {
	errs := []error{}

	for _, validation := range this.validations {
		if err := validation(v); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}

func (this *stringValidations[T]) Min(n int) *stringValidations[T] {
	this.validations = append(this.validations, StringMinFn[T](this.name, n))
	return this
}

func (this *stringValidations[T]) Max(n int) *stringValidations[T] {
	this.validations = append(this.validations, StringMaxFn[T](this.name, n))
	return this
}

func (this *stringValidations[T]) Regex(pattern string) *stringValidations[T] {
	this.validations = append(this.validations, StringRegexFn[T](this.name, pattern))
	return this
}

func (this *stringValidations[T]) NotEmpty(pattern string) *stringValidations[T] {
	this.validations = append(this.validations, StringNotEmptyFn[T](this.name))
	return this
}

func (this *stringValidations[T]) UpperCase(pattern string) *stringValidations[T] {
	this.validations = append(this.validations, StringUpperCaseFn[T](this.name))
	return this
}

func (this *stringValidations[T]) LowerCase(pattern string) *stringValidations[T] {
	this.validations = append(this.validations, StringLowerCaseFn[T](this.name))
	return this
}

func (this *stringValidations[T]) Required() *stringValidations[T] {
	this.validations = append(this.validations, StringRequiredFn[T](this.name))
	return this
}
