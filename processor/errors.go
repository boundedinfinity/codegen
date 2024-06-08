package processor

import (
	"boundedinfinity/codegen/model"
	"errors"
)

var (
	ErrCodeGenSourceDuplicate   = errors.New("duplicate source")
	ErrCodeGenSourceDuplicateFn = func(path string) error {
		return errors.Join(errors.New(path), ErrCodeGenSourceDuplicate)
	}

	ErrCodeGenTypeSchemaIdDuplicate   = errors.New("duplicate schema-id")
	ErrCodeGenTypeSchemaIdDuplicateFn = func(obj model.CodeGenType) error {
		return errors.Join(errors.New(obj.TypeId().Get()), ErrCodeGenTypeSchemaIdDuplicate)
	}

	ErrCodeGenCantReadFile   = errors.New("can't read file")
	ErrCodeGenCantReadFileFn = func(path string, err error) error {
		return errors.Join(errors.New(path), ErrCodeGenCantReadFile)
	}

	ErrCodeGenUnsupportedFileType   = errors.New("unsupported file type")
	ErrCodeGenUnsupportedFileTypeFn = func(ext string) error {
		return errors.Join(errors.New(ext), ErrCodeGenUnsupportedFileType)
	}

	ErrCodeGenUnsupportedType   = errors.New("unsupported file type")
	ErrCodeGenUnsupportedTypeFn = func(typ model.CodeGenType) error {
		return errors.Join(errors.New(typ.CodeGenId()), ErrCodeGenUnsupportedType)
	}
)
