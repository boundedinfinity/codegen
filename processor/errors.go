package processor

import (
	"boundedinfinity/codegen/model"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrCodeGenSourceDuplicate   = errorer.New("duplicate source")
	ErrCodeGenSourceDuplicateFn = ErrCodeGenSourceDuplicate.ValueFn()

	ErrCodeGenTypeSchemaIdDuplicate   = errorer.New("duplicate schema-id")
	ErrCodeGenTypeSchemaIdDuplicateFn = func(obj model.CodeGenSchema) error {
		return ErrCodeGenTypeSchemaIdDuplicate.WithValue(obj.Common().Id.Get())
	}

	ErrCodeGenCantReadFile   = errorer.New("can't read file")
	ErrCodeGenCantReadFileFn = func(path string, err error) error {
		return ErrCodeGenUnsupportedFileType.Sub(err).WithValue(path)
	}

	ErrCodeGenUnsupportedFileType   = errorer.New("unsupported file type")
	ErrCodeGenUnsupportedFileTypeFn = ErrCodeGenUnsupportedFileType.ValueFn()

	ErrCodeGenUnsupportedType   = errorer.New("unsupported file type")
	ErrCodeGenUnsupportedTypeFn = func(typ model.CodeGenSchema) error {
		return ErrCodeGenUnsupportedType.WithValue(typ.Schema())
	}
)
