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
	ErrUnsupportedScheme  = errors.New("unsupported scheme")
	ErrMissingName        = errors.New("missing name")
	ErrInvalidSchemaType  = errors.New("invalid schema type")
)

func ErrDuplicateSourceUriV(v string) error { return nE(ErrDuplicateSourceUri, v) }
func ErrCodeGenIdDuplicateV(v string) error { return nE(ErrCodeGenIdDuplicate, v) }
func ErrTemplateDuplicateV(v string) error  { return nE(ErrTemplateDuplicate, v) }
func ErrUnsupportedSchemeV(v string) error  { return nE(ErrUnsupportedScheme, v) }

func nE(err error, v interface{}) error {
	return fmt.Errorf("%w : %v", err, v)
}
