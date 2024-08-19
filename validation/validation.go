// Package validation stuff goes here
package validation

import (
	"errors"
	"fmt"
)

type Validater[T any] interface {
	Validate(T) error
}

type Validation[T any] func(T) error

var ErrValidatorIncorrectType = errors.New("incorrect type")

type ErrValidatorIncorrectTypeDetails struct {
	name  string
	value any
}

func (this ErrValidatorIncorrectTypeDetails) Error() string {
	return fmt.Sprintf("%s for %s: %v", ErrValidatorIncorrectType, this.name, this.value)
}

func (this ErrValidatorIncorrectTypeDetails) Unwrap() error {
	return ErrValidatorIncorrectType
}

type validator[T any] struct {
	name        string
	validations []Validation[T]
}

func (this *validator[T]) append(validation Validation[T]) {
	this.validations = append(this.validations, validation)
}

func (this validator[T]) Validate(value any) error {
	tvalue, ok := value.(T)

	if !ok {
		return &ErrValidatorIncorrectTypeDetails{name: this.name, value: value}
	}

	for _, validation := range this.validations {
		if err := validation(tvalue); err != nil {
			return err
		}
	}

	return nil
}

// Indexed Name
func in(name string, length, index int) string {
	if length > 0 {
		return fmt.Sprintf("%s[%d]", name, index)
	}

	return name
}
