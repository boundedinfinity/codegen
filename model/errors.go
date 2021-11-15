package model

import (
	"errors"
	"fmt"
)

var (
	ErrDuplicateSourceUri = errors.New("duplicate uri")
	ErrCodeGenIdEmpty     = errors.New("empty codegen schema ID")
	ErrCodeGenIdDuplicate = errors.New("duplicate codegen schema ID")
)

func ErrDuplicateSourceUriV(v string) error {
	return fmt.Errorf("%w : %v", ErrDuplicateSourceUri, v)
}

func ErrCodeGenIdDuplicateV(v string) error {
	return fmt.Errorf("%w : %v", ErrCodeGenIdDuplicate, v)
}
