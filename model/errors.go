package model

import "github.com/boundedinfinity/go-commoner/errorer"

var (
	ErrDuplicateSourceUri                = errorer.Errorf("duplicate uri")
	ErrDuplicateSourceUriv               = errorer.Errorfn(ErrDuplicateSourceUri)
	ErrCodeGenIdEmpty                    = errorer.Errorf("empty codegen schema ID")
	ErrCodeGenIdDuplicate                = errorer.Errorf("duplicate codegen schema ID")
	ErrCodeGenIdDuplicatev               = errorer.Errorfn(ErrCodeGenIdDuplicate)
	ErrTemplateEmpty                     = errorer.Errorf("duplicate empty")
	ErrTemplateDuplicate                 = errorer.Errorf("duplicate template")
	ErrTemplateDuplicatev                = errorer.Errorfn(ErrTemplateDuplicate)
	ErrUnsupportedScheme                 = errorer.Errorf("unsupported scheme")
	ErrUnsupportedSchemev                = errorer.Errorfn(ErrUnsupportedScheme)
	ErrMissingName                       = errorer.Errorf("missing name")
	ErrInvalidSchemaType                 = errorer.Errorf("invalid schema type")
	ErrPathDuplicate                     = errorer.Errorf("code gen schema path duplicate")
	ErrPathDuplicatev                    = errorer.Errorfn(ErrPathDuplicate)
	ErrMimeTypeUnsupported               = errorer.Errorf("MIME type unsupported")
	ErrMimeTypeUnsupportedv              = errorer.Errorfn(ErrMimeTypeUnsupported)
	ErrCodeGenRefNotFound                = errorer.Errorf("code gen ref not found")
	ErrCodeGenRefNotFoundv               = errorer.Errorfn(ErrCodeGenRefNotFound)
	ErrCodeGenOperationDuplicate         = errorer.Errorf("duplicate codegen schema operation")
	ErrCodeGenOperationDuplicatev        = errorer.Errorfn(ErrCodeGenOperationDuplicate)
	ErrCodeGenMappingsPackageDuplicate   = errorer.Errorf("codegen.mappings.package already defined")
	ErrCodeGenMappingsPackageDuplicatev  = errorer.Errorfn(ErrCodeGenMappingsPackageDuplicate)
	ErrCodeGenMappingsRootDirDuplicate   = errorer.Errorf("duplicate codegen.mappings.rootDir")
	ErrCodeGenMappingsRootDirDuplicatev  = errorer.Errorfn(ErrCodeGenMappingsRootDirDuplicate)
	ErrCodeGenMappingsRelpaceDuplicate   = errorer.Errorf("duplicate codegen.mappings.replace")
	ErrCodeGenMappingsRelpaceDuplicatev  = errorer.Errorfn(ErrCodeGenMappingsRelpaceDuplicate)
	ErrCodeGenTemplateFilePathDuplicate  = errorer.Errorf("duplicate codegen.templates.file.path")
	ErrCodeGenTemplateFilePathDuplicatev = errorer.Errorfn(ErrCodeGenTemplateFilePathDuplicate)
	ErrMinMax                            = errorer.Errorf("max is less than min")
)
