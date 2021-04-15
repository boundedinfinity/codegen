package generator

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/optional"
	"github.com/ozgio/strutil"
)

func ifeq(v1, v2 string) bool {
	return v1 == v2
}

func uc(v string) string {
	return strings.ToUpper(v)
}

func ucFirst(v string) string {
	f := string(v[0])
	r := string(v[1:])
	return strings.ToUpper(f) + r
}

func camelCase(v string) string {
	return lcFirst(strutil.ToCamelCase(v))
}

func lc(v string) string {
	return strings.ToLower(v)
}

func lcFirst(v string) string {
	f := string(v[0])
	r := string(v[1:])
	return strings.ToLower(f) + r
}

func basePath(v string) string {
	return path.Base(v)
}

func filename(v string) string {
	var fn string
	fn = v
	fn = filepath.Base(fn)
	fn = strings.TrimSuffix(fn, filepath.Ext(fn))
	return fn
}

func peq(v1 optional.StringOptional, v2 string) bool {
	if v1.IsEmpty() {
		return false
	}

	return v1.Get() == v2
}
