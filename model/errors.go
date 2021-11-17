package model

import (
	"errors"
	"fmt"
)

var (
	ErrDuplicateSourceUri = errors.New("duplicate uri")
	ErrCodeGenIdEmpty     = errors.New("empty codegen schema ID")
	ErrCodeGenIdDuplicate = errors.New("duplicate codegen schema ID")
	ErrTemplateEmpty      = errors.New("duplicate empty")
	ErrTemplateDuplicate  = errors.New("duplicate template")
)

func ErrDuplicateSourceUriV(v string) error {
	return fmt.Errorf("%w : %v", ErrDuplicateSourceUri, v)
}

func ErrCodeGenIdDuplicateV(v string) error {
	return fmt.Errorf("%w : %v", ErrCodeGenIdDuplicate, v)
}

func ErrTemplateDuplicateV(v string) error {
	return fmt.Errorf("%w : %v", ErrTemplateDuplicate, v)
}
