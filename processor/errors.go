package processor

import (
	"boundedinfinity/codegen/model"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrCodeGenSourceDuplicate   = errorer.New("duplicate source")
	ErrCodeGenSourceDuplicateFn = ErrCodeGenSourceDuplicate.ValueFn()

	ErrCodeGenTypeSchemaIdDuplicate   = errorer.New("duplicate schema-id")
	ErrCodeGenTypeSchemaIdDuplicateFn = func(obj model.CodeGenType) error {
		return ErrCodeGenTypeSchemaIdDuplicate.WithValue(obj.QName().Get())
	}

	ErrCodeGenCantReadFile   = errorer.New("can't read file")
	ErrCodeGenCantReadFileFn = func(path string, err error) error {
		return ErrCodeGenUnsupportedFileType.Sub(err).WithValue(path)
	}

	ErrCodeGenUnsupportedFileType   = errorer.New("unsupported file type")
	ErrCodeGenUnsupportedFileTypeFn = ErrCodeGenUnsupportedFileType.ValueFn()

	ErrCodeGenUnsupportedType   = errorer.New("unsupported file type")
	ErrCodeGenUnsupportedTypeFn = func(typ model.CodeGenType) error {
		return ErrCodeGenUnsupportedType.WithValue(typ.BaseType())
	}
)
