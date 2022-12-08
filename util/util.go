package util

import (
	"boundedinfinity/codegen/canonical/canonical_type"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/extentioner"
	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-commoner/trier"
	"github.com/boundedinfinity/go-jsonschema/schematype"
	"github.com/boundedinfinity/go-mimetyper/file_extention"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

var (
	codegenExts = []string{
		".codegen.json",
		".codegen.yaml",
		".codegen.yml",
	}

	codegenTypeExts = []string{
		".codegen-type.json",
		".codegen-type.yaml",
		".codegen-type.yml",
	}

	jsonSchemaExts = []string{
		".json-schema.json",
		".json-schema.yaml",
		".json-schema.yml",
	}
)

func IsCodeGenSchemaTypeFile(v string) bool {
	return slicer.ContainsFn(codegenTypeExts, func(x string) bool {
		return strings.HasSuffix(v, x)
	})
}

func IsCodeGenSchemaFile(v string) bool {
	return slicer.ContainsFn(codegenExts, func(x string) bool {
		return strings.HasSuffix(v, x)
	})
}

func IsJsonSchemaFile(v string) bool {
	return slicer.ContainsFn(jsonSchemaExts, func(x string) bool {
		return strings.HasSuffix(v, x)
	})
}

func IsSchemaFile(v string) bool {
	return IsCodeGenSchemaFile(v) || IsJsonSchemaFile(v) || IsCodeGenSchemaTypeFile(v)
}

func IsJsonSchemaTemplate(typ optioner.Option[schematype.SchemaType], path string) bool {
	if typ.Empty() {
		return false
	}

	base := filepath.Base(path)
	ts := typ.String()
	return strings.HasPrefix(base, ts)
}

func GetCanonicalType(path string) optioner.Option[canonical_type.CanonicalType] {
	filename := pather.Base(path)
	found, ok := slicer.FindFn(canonical_type.All, func(v canonical_type.CanonicalType) bool {
		return strings.HasPrefix(filename, string(v))
	})

	if ok {
		return optioner.Some(found)
	}

	return optioner.None[canonical_type.CanonicalType]()
}

func GetTemplateType(path string) trier.Try[mime_type.MimeType] {
	var ext string
	ext = extentioner.Ext(path)
	tm, err := file_extention.GetMimeType(ext)
	return trier.Complete(tm, err)
}

func GetOutputType(path string) trier.Try[mime_type.MimeType] {
	ext := path
	ext = extentioner.Strip(ext)
	ext = extentioner.Ext(ext)
	tm, err := file_extention.GetMimeType(ext)
	return trier.Complete(tm, err)
}
