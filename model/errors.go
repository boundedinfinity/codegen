package model

import "github.com/boundedinfinity/go-commoner/errorer"

var (
	ErrDuplicateSourceUri               = errorer.Errorf("duplicate uri")
	ErrCodeGenIdEmpty                   = errorer.Errorf("empty codegen schema ID")
	ErrCodeGenIdDuplicate               = errorer.Errorf("duplicate codegen schema ID")
	ErrTemplateEmpty                    = errorer.Errorf("duplicate empty")
	ErrTemplateDuplicate                = errorer.Errorf("duplicate template")
	ErrUnsupportedScheme                = errorer.Errorf("unsupported scheme")
	ErrMissingName                      = errorer.Errorf("missing name")
	ErrInvalidSchemaType                = errorer.Errorf("invalid schema type")
	ErrPathDuplicate                    = errorer.Errorf("code gen schema path duplicate")
	ErrMimeTypeUnsupported              = errorer.Errorf("MIME type unsupported")
	ErrCodeGenRefNotFound               = errorer.Errorf("code gen ref not found")
	ErrCodeGenOperationDuplicate        = errorer.Errorf("duplicate codegen schema operation")
	ErrCodeGenMappingsPackageDuplicate  = errorer.Errorf("codegen.mappings.package already defined")
	ErrCodeGenMappingsRootDirDuplicate  = errorer.Errorf("duplicate codegen.mappings.rootDir")
	ErrCodeGenMappingsRelpaceDuplicate  = errorer.Errorf("duplicate codegen.mappings.replace")
	ErrCodeGenTemplateFilePathDuplicate = errorer.Errorf("duplicate codegen.templates.file.path")
	ErrMinMax                           = errorer.Errorf("max is less than min")
	ErrProjectPackageAlreadyDefined     = errorer.New("project package already defined")
	ErrRefNotFound                      = errorer.New("reference not found")
	ErrRefEmpty                         = errorer.New("empty reference")
)
