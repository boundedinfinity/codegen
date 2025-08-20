package errorer

import (
	"errors"
	"fmt"
)

type Errorer struct {
	err error
}

func (this Errorer) Error() string {
	return this.err.Error()
}

func New(text string) error {
	return &Errorer{
		err: errors.New(text),
	}
}

func ValueFn(err error) func(...any) error {
	return ValueFnf(err, "%v")
}

func ValueFnf(err error, format string) func(...any) error {
	if format == "" {
		format = "%v"
	}

	return func(a ...any) error {
		aa := append([]any{err}, a...)
		return fmt.Errorf("%w : "+format, aa...)
	}
}
