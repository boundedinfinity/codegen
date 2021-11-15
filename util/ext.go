package util

import (
	"strings"
)

func IsCodeGenFile(v string) bool {
	var x bool

	if strings.HasSuffix(v, ".codegen.json") {
		x = true
	}

	if strings.HasSuffix(v, ".codegen.yaml") {
		x = true
	}

	if strings.HasSuffix(v, ".codegen.yml") {
		x = true
	}

	return x
}

func IsJsonSchemaFile(v string) bool {
	var x bool

	if strings.HasSuffix(v, ".schema.json") {
		x = true
	}

	if strings.HasSuffix(v, ".schema.yaml") {
		x = true
	}

	if strings.HasSuffix(v, ".schema.yml") {
		x = true
	}

	return x
}
