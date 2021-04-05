package util

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"path"
	"strings"
)

func ifeq(v1, v2 string) bool {
	return v1 == v2
}

func uc(v string) string {
	return strings.ToUpper(v)
}

func ucFirst(v string) string {
	return strings.Title(v)
}

func lc(v string) string {
	return strings.ToLower(v)
}

func lcFirst(v string) string {
	return strings.Title(v)
}

func basePath(v string) string {
	return path.Base(v)
}

func peq(v1 *string, v2 string) bool {
	if v1 == nil {
		return false
	}

	return *v1 == v2
}

func typeGo(v model.JsonSchema_Draft07) string {
	if v.Ref != nil {
		return ucFirst(path.Base(*v.Ref))
	} else {
		switch *v.Type {
		case "string":
			return "string"
		case "integer":
			return "int64"
		case "number":
			return "float32"
		case "boolean":
			return "bool"
		case "array":
			if v.Items != nil && v.Items.Type != nil {
				return fmt.Sprintf("[]%v", *v.Items.Type)
			} else if v.Items != nil && v.Items.Ref != nil {
				return fmt.Sprintf("[]%v", ucFirst(path.Base(*v.Items.Ref)))
			}
			return "array"
		case "null":
			return "nil"
		default:
			return "unknown-type"
		}
	}
}
