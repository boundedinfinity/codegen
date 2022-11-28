package util

import (
	"strings"

	"github.com/boundedinfinity/go-commoner/slicer"
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
