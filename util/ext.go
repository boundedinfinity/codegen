package util

import (
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-jsonschema/schematype"
)

var (
	codegenExts = []string{
		".codegen.json",
		".codegen.yaml",
		".codegen.yml",
	}

	jsonSchemaExts = []string{
		".schema.json",
		".schema.yaml",
		".schema.yml",
	}
)

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
	return IsCodeGenSchemaFile(v) || IsJsonSchemaFile(v)
}

func IsJsonSchemaTemplate(typ optioner.Option[schematype.SchemaType], path string) bool {
	if typ.Empty() {
		return false
	}

	base := filepath.Base(path)
	ts := typ.String()
	return strings.HasPrefix(base, ts)
}
