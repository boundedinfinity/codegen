package errutil

import (
	"fmt"
	"strings"
)

type CodeGenError struct {
	path    []string
	message string
}

func (t CodeGenError) Error() string {
	var s string

	if t.path == nil && len(t.path) > 0 {
		s = strings.Join(t.path, ".")
		s = fmt.Sprintf("%v : ", s)
	}

	s = fmt.Sprintf("%v%v", s, t.message)

	return s
}

func Errorf(path []string, format string, a ...interface{}) error {
	return CodeGenError{
		path:    path,
		message: fmt.Sprintf(format, a...),
	}
}

func CannotBeEmpty(path []string) error {
	return Errorf(path, "cannot be empty")
}

func NotFound(path []string) error {
	return Errorf(path, "not found")
}

func InvalidateType(path []string) error {
	return Errorf(path, "invalid type")
}

func MustBeOneOf(path []string, oneOf []string) error {
	var message string

	message = strings.Join(oneOf, ",")
	message = fmt.Sprintf("invalid type, must be one of: %v", message)

	return Errorf(path, message)
}
