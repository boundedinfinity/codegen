package validation

type stringValidations[T ~string] struct {
	validator *validator[T]
}

var _ Validater[string] = &stringValidations[string]{}

func String[T ~string](name string) *stringValidations[T] {
	return &stringValidations[T]{validator: &validator[T]{name: name}}
}

func (this *stringValidations[T]) Validate(value T) error {
	return this.validator.Validate(value)
}

func (this *stringValidations[T]) Min(n int) *stringValidations[T] {
	this.validator.append(StringMinFn[T](this.validator.name, n))
	return this
}

func (this *stringValidations[T]) Max(n int) *stringValidations[T] {
	this.validator.append(StringMaxFn[T](this.validator.name, n))
	return this
}

func (this *stringValidations[T]) Regex(pattern string) *stringValidations[T] {
	this.validator.append(StringRegexFn[T](this.validator.name, pattern))
	return this
}

func (this *stringValidations[T]) NotEmpty() *stringValidations[T] {
	this.validator.append(StringNotEmptyFn[T](this.validator.name))
	return this
}

func (this *stringValidations[T]) UpperCase() *stringValidations[T] {
	this.validator.append(StringUpperCaseFn[T](this.validator.name))
	return this
}

func (this *stringValidations[T]) LowerCase(pattern string) *stringValidations[T] {
	this.validator.append(StringLowerCaseFn[T](this.validator.name))
	return this
}

func (this *stringValidations[T]) Required() *stringValidations[T] {
	this.validator.append(StringRequiredFn[T](this.validator.name))
	return this
}
