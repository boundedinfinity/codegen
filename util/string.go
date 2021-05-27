package util

import (
	"boundedinfinity/codegen/model"
	"encoding/json"
	"path"
	"strings"

	"github.com/boundedinfinity/caser"
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
	return caser.PhraseToCamel(v)
}

func PascalCase(v string) string {
	return caser.PhraseToCamel(v)
}

func PathBase(v string) string {
	return path.Base(v)
}

func PathDir(v string) string {
	return path.Dir(v)
}

func PathDirBase(v string) string {
	o := v

	if strings.Contains(v, "/") {
		o = path.Base(path.Dir(v))
	} else {
		o = v
	}

	return o
}

func SameNamespace(a, b model.OutputModel) bool {
	aNs := path.Dir(a.FullName)
	bNs := path.Dir(b.FullName)
	return aNs == bNs
}

func ToJson(v interface{}) string {
	j, _ := json.MarshalIndent(v, "", "    ")
	return string(j)
}
