package util

import (
	"boundedinfinity/codegen/model"
	"path"
	"strings"

	"github.com/ozgio/strutil"
)

func IsSchemaInt(v *model.OutputModel) bool {
	return v.Type == model.SchemaType_Int
}

func IsSchemaString(v *model.OutputModel) bool {
	return v.Type == model.SchemaType_String
}

func Uc(v string) string {
	return strings.ToUpper(v)
}

func UcFirst(v string) string {
	f := string(v[0])
	r := string(v[1:])
	return strings.ToUpper(f) + r
}

func Lc(v string) string {
	return strings.ToLower(v)
}

func LcFirst(v string) string {
	f := string(v[0])
	r := string(v[1:])
	return strings.ToLower(f) + r
}

func CamelCase(v string) string {
	return LcFirst(strutil.ToCamelCase(v))
}

func NsBase(v string) string {
	return path.Base(v)
}

func NsDir(v string) string {
	return path.Dir(v)
}

func SameNamespace(a, b model.OutputModel) bool {
	return path.Dir(a.Name) == path.Dir(b.Name)
}
