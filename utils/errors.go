package utils

import "fmt"

func ErrorValue[T any](parent error) errorFactory[T] {
	return ErrorFactory(parent, func(value T) string {
		return fmt.Sprintf("%v", value)
	})
}

func ErrorFactory[T any](parent error, fn func(T) string) errorFactory[T] {
	return errorFactory[T]{
		fn:     fn,
		parent: parent,
	}
}

type errorFactory[T any] struct {
	parent error
	fn     func(T) string
}

func (this errorFactory[T]) New(item T) error {
	return &factoryError{
		parents: this.parent,
		new:     fmt.Errorf("%w : %s", this.parent, this.fn(item)),
	}
}

type factoryError struct {
	parents error
	new     error
}

func (this *factoryError) Error() string {
	return this.new.Error()
}

func (this *factoryError) Unwrap() []error {
	return []error{
		this.parents,
		this.new,
	}
}
